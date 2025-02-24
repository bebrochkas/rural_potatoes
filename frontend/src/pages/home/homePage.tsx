import { axiosInst } from "@/api/axios";
import { FilmCard } from "@/components/films/card";
import { TagBar } from "@/components/tags/tagBar";
import { Film, Tag } from "@/interfaces/interfaces";
import { useEffect, useState, useCallback, useRef } from "react";

const PAGE_SIZE = 4; // Number of films per request

export const HomePage = () => {
    const [films, setFilms] = useState<Film[]>([]);
    const [offset, setOffset] = useState(0);
    const [hasMore, setHasMore] = useState(true);
    const [isLoading, setIsLoading] = useState<boolean>(false); // Prevent multiple fetches at once

    const [tags, setTags] = useState<Tag[]>([]);
    const [selectedTags, setSelectedTags] = useState<number[]>([]);

    useEffect(() => {
        const fetchTags = async () => {
            try {
                const response = await axiosInst.get<Tag[]>(`tags/`);
                setTags((prevTags) => [...prevTags, ...response.data]); // Append new films
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
    }, [selectedTags]);

    const fetchProducts = useCallback(async () => {
        if (!hasMore) {
            return;
        }
        setIsLoading(true);
        try {
            setOffset((prevOffset) => prevOffset + PAGE_SIZE);

            const response = await axiosInst.get<Film[]>(
                `films/fetch?offset=${offset}&limit=${PAGE_SIZE}&tags=${selectedTags}`,
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
    }, [offset, hasMore, selectedTags]);

    const observerTarget = useRef(null);

    useEffect(() => {
        const observer = new IntersectionObserver(
            (entries) => {
                if (entries[0].isIntersecting && hasMore) {
                    fetchProducts();
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
    }, [observerTarget, fetchProducts, hasMore]);

    return (
        <div className="flex flex-col gap-4">
            <TagBar
                tags={tags}
                selectedTags={selectedTags}
                setSelectedTags={setSelectedTags}
            />
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
