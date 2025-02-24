import { Badge } from "@/components/ui/badge";

import { Tag } from "@/interfaces/interfaces";

interface tagBarProps {
    tags: Tag[];
    selectedTags: number[];
    setSelectedTags: React.Dispatch<React.SetStateAction<number[]>>;
}

export const TagBar = ({
    tags,
    selectedTags,
    setSelectedTags,
}: tagBarProps) => {
    return (
        <div className="flex w-full flex-col overflow-x-scroll">
            <ul className="flex flex-row gap-2">
                {tags.map((tag) => (
                    <li key={tag.id}>
                        <Badge
                            className={`rounded-full text-nowrap font-medium cursor-pointer ${
                                selectedTags.includes(tag.id)
                                    ? "bg-blue-500 text-white"
                                    : "bg-gray-200"
                            }`}
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
