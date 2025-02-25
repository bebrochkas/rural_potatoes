import { axiosInst } from "@/api/axios";
import { FilmCard } from "@/components/films/card";
import { Input } from "@/components/ui/input";
import { TagBar } from "@/components/tags/tagBar";
import { Film, Tag } from "@/interfaces/interfaces";
import { useEffect, useState, useCallback, useRef } from "react";
import { Badge } from "@/components/ui/badge";

import { debounce } from "lodash";

const PAGE_SIZE = 4; // Number of films per request

export const HomePage = () => {
    const [films, setFilms] = useState<Film[]>([]);
    const [offset, setOffset] = useState(0);
    const [hasMore, setHasMore] = useState(true);
    const [isLoading, setIsLoading] = useState<boolean>(false); // Prevent multiple fetches at once

    const [feedMode, setFeedMode] = useState<boolean>(false);
    const [prompt, setPrompt] = useState<string>("");
    const promptInpRef = useRef<HTMLInputElement>(null);
    const [tagSuggest, setTagSuggest] = useState<Tag[]>([]);

    const [tags, setTags] = useState<Tag[]>([]);
    const [selectedTags, setSelectedTags] = useState<number[]>([]);

    useEffect(() => {
        const fetchTags = async () => {
            try {
                const response = await axiosInst.get<Tag[]>(`tags/`);
                setTags(response.data); // Append new films
            } catch (error) {
                console.log((error as Error).message);
            }
        };

        fetchTags();
    }, []);

    useEffect(() => {
        setTags((prevTags) => {
            const nonSelected = prevTags.filter(
                (tag) => !selectedTags.includes(tag.id),
            );
            const selected = prevTags.filter((tag) =>
                selectedTags.includes(tag.id),
            );
            return [...selected, ...nonSelected];
        });

        setOffset(0);
        setFilms([]);
        setHasMore(true);
    }, [selectedTags, feedMode]);

    const fetchSuggestion = async (prompt: string) => {
        if (prompt) {
            setTagSuggest(
                tags.filter((tag) =>
                    tag.name.toLowerCase().startsWith(prompt.toLowerCase()),
                ),
            );

            if (prompt.length >= 2 && prompt.endsWith(" ")) {
                try {
                    const response = await axiosInst.post<Tag[]>(
                        `tags/suggest?q=${prompt}`,
                    );

                    setTagSuggest(response.data);
                } catch (error) {
                    console.error("Error fetching images:", error);
                }
            }
        } else {
            setTagSuggest([]);
        }
    };

    const debouncedSearch = debounce(async (value: string) => {
        setPrompt(value);
        setOffset(0);
        setFilms([]);
        setHasMore(true);
    }, 100);

    const fetchFilms = useCallback(async () => {
        if (!hasMore) {
            return;
        }
        setIsLoading(true);
        try {
            setOffset((prevOffset) => prevOffset + PAGE_SIZE);

            const response = await axiosInst.get<Film[]>(
                `films?offset=${offset}&limit=${PAGE_SIZE}&tags=${feedMode ? "feed" : selectedTags}&q=${prompt}`,
            );

            const newFilms = response.data;

            if (newFilms.length < PAGE_SIZE) {
                setHasMore(false);
            }

            setFilms((prevFilms) => [...prevFilms, ...response.data]); // Append new films
        } catch (error) {
            console.error("Error fetching images:", error);
            setHasMore(false);
            setIsLoading(false);
        } finally {
            setIsLoading(false);
        }
    }, [offset, hasMore, selectedTags, prompt, feedMode]);

    const observerTarget = useRef(null);

    useEffect(() => {
        const observer = new IntersectionObserver(
            (entries) => {
                if (entries[0].isIntersecting && hasMore) {
                    fetchFilms();
                }
            },
            { threshold: 1 },
        );

        if (observerTarget.current) {
            observer.observe(observerTarget.current);
        }

        return () => {
            if (observerTarget.current) {
                observer.unobserve(observerTarget.current);
            }
        };
    }, [observerTarget, fetchFilms, hasMore]);

    return (
        <div className="flex flex-col gap-10 justify-center items-center">
            <div className="flex flex-col gap-1 w-[45%]">
                <div className="flex flex-row">
                    <Input
                        ref={promptInpRef}
                        placeholder="Начните писать своими словами или сразу введите конкретное название"
                        onChange={(e) => {
                            fetchSuggestion(e.target.value);
                            debouncedSearch(e.target.value);
                        }}
                    />

                    {tagSuggest.map((tag) => (
                        <Badge
                            className={
                                "rounded-full text-nowrap font-medium cursor-pointer "
                            }
                            onClick={() => {
                                setSelectedTags((prevTags) => {
                                    if (prevTags.includes(tag.id)) {
                                        return prevTags.filter(
                                            (id) => id !== tag.id,
                                        );
                                    } else {
                                        setPrompt("");
                                        promptInpRef.current!.value = "";
                                        return [...prevTags, tag.id];
                                    }
                                });
                                setTagSuggest((prev) => {
                                    return prev.filter(
                                        (suggestTag) => suggestTag.id != tag.id,
                                    );
                                });
                            }}
                        >
                            {tag.name}
                        </Badge>
                    ))}
                </div>

                {prompt && (
                    <blockquote className="mt-2 border-l-2 pl-2 italic">
                        Нажмите пробел после ввода чтобы получить AI подсказки
                        по тегам
                    </blockquote>
                )}
                <TagBar
                    tags={tags}
                    selectedTags={selectedTags}
                    setSelectedTags={setSelectedTags}
                    feedMode={feedMode}
                    setFeedMode={setFeedMode}
                />
            </div>

            <ul className="grid items-center gap-8 md:px-6 lg:gap-12 grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {films.map((film) => (
                    <li key={film.id}>
                        <FilmCard film={film} />
                    </li>
                ))}
            </ul>
            {isLoading && <p>Loading...</p>}
            {!hasMore && <p className="text-center">the end.</p>}

            <div ref={observerTarget}></div>
        </div>
    );
};
