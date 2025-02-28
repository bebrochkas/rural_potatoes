import { axiosInst } from "@/api/axios";
import { FilmCard } from "@/components/films/card";
import { Input } from "@/components/ui/input";
import { TagBar } from "@/components/tags/tagBar";
import { Film, Tag } from "@/interfaces/interfaces";
import { useEffect, useState, useCallback, useRef } from "react";
import { Badge } from "@/components/ui/badge";
import shinjiUrl from "@/assets/shinj.webp";
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
        setIsLoading(true);

        setHasMore(true);
    }, [selectedTags, feedMode]);

    const fetchSuggestion = async (prompt: string) => {
        if (prompt.length > 2) {
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
        setIsLoading(true);
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
    }, [observerTarget, fetchFilms, hasMore, feedMode]);

    return (
        <div className="flex flex-col gap-10 justify-center items-center ">
            <div className="flex flex-col gap-2 w-full lg:w-[45%] md:w-3/4 ">
                <a
                    className={`italic text-sm text-center text-gray-500  transition-opacity duration-500 ease-in-out ${
                        prompt ? "opacity-100" : "opacity-0"
                    }`}
                >
                    Нажмите пробел после ввода чтобы получить AI подсказки по
                    тегам
                </a>

                <div className="flex flex-row  justify-center">
                    <Input
                        className={`rounded-full transition-all duration-500 ease-in-out ${prompt.length !== 0 ? "w-full" : "w-3/4"}`}
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
                                " rounded-full  text-nowrap font-medium cursor-pointer "
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

                <TagBar
                    tags={tags}
                    selectedTags={selectedTags}
                    setSelectedTags={setSelectedTags}
                    feedMode={feedMode}
                    setFeedMode={setFeedMode}
                />
            </div>

            <ul className="grid items-center gap-8 md:px-6 lg:gap-12 grid-cols-2 md:grid-cols-3 lg:grid-cols-4  ">
                {films.map((film) => (
                    <li key={film.id}>
                        <FilmCard film={film} />
                    </li>
                ))}
            </ul>

            {isLoading && (
                <div
                    role="status"
                    className="w-full h-full flex items-center justify-center"
                >
                    <svg
                        aria-hidden="true"
                        className="w-8 h-8 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600"
                        viewBox="0 0 100 101"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <path
                            d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                            fill="currentColor"
                        />
                        <path
                            d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                            fill="currentFill"
                        />
                    </svg>
                    <span className="sr-only">Loading...</span>
                </div>
            )}
            {!hasMore && (
                <div className="w-full flex flex-col items-center justify-center gap-2 font-mono">
                    На этом фильмы кончаются.
                    <img className="" src={shinjiUrl}></img>
                </div>
            )}

            <div ref={observerTarget}></div>
        </div>
    );
};
