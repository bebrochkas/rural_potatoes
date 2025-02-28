import { Badge } from "@/components/ui/badge";

import { Tag } from "@/interfaces/interfaces";

interface tagBarProps {
    tags: Tag[];
    selectedTags: number[];
    setSelectedTags: React.Dispatch<React.SetStateAction<number[]>>;
    feedMode: boolean;
    setFeedMode: React.Dispatch<React.SetStateAction<boolean>>;
}

export const TagBar = ({
    tags,
    selectedTags,
    setSelectedTags,
    feedMode,
    setFeedMode,
}: tagBarProps) => {
    return (
        <div className="flex w-full flex-col overflow-x-scroll no-scrollbar rounded-full ">
            <ul className="flex flex-row gap-2 items-center rounded-full">
                <li>
                    <Badge
                        className={`rounded-full text-nowrap font-medium cursor-pointer text-sm ${
                            feedMode ? "bg-blue-700 text-white" : "bg-gray-400"
                        }`}
                        onClick={() => {
                            setFeedMode((prev) => !prev);
                        }}
                    >
                        Рекомедации
                    </Badge>
                </li>

                {tags.map((tag) => (
                    <li key={tag.id}>
                        <Badge
                            className={`rounded-full text-nowrap font-medium text-sm cursor-pointer`}
                            style={{
                                backgroundColor: `${
                                    selectedTags.includes(tag.id)
                                        ? "var(--color-blue-500)"
                                        : tag.hex
                                }`,
                            }}
                            onClick={() => {
                                setSelectedTags((prevTags) => {
                                    if (prevTags.includes(tag.id)) {
                                        return prevTags.filter(
                                            (id) => id !== tag.id,
                                        );
                                    } else {
                                        return [...prevTags, tag.id];
                                    }
                                });
                            }}
                        >
                            {tag.name}
                        </Badge>
                    </li>
                ))}
            </ul>
        </div>
    );
};
