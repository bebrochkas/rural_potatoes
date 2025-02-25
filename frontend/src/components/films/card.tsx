import { axiosInst } from "@/api/axios";
import { Badge } from "@/components/ui/badge";
import { Film } from "@/interfaces/interfaces";

import { ThumbsUp, ThumbsDown } from "lucide-react";
import { Button } from "@/components/ui/button";

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";

export const FilmCard = ({ film }: { film: Film }) => {
    const rate = async (action: string) => {
        try {
            const _ = await axiosInst.post(
                `films/rate?id=${film.id}&action=${action}`,
            );
        } catch (error) {
            console.log((error as Error).message);
        }
    };

    const Thumbs = ({ hoverCard }: { hoverCard: boolean }) => {
        return (
            <div className="flex gap-1 ">
                <Button
                    variant={hoverCard ? "outline" : "ghost"}
                    className="rounded-full border-2 border-gray-300"
                    onClick={() => rate("like")}
                >
                    {hoverCard ? "" : "Нравиться"}
                    <a>{film.likes}</a>
                    <ThumbsUp strokeWidth={2} />
                </Button>
                <Button
                    variant={hoverCard ? "secondary" : "ghost"}
                    className="rounded-full border-2 border-gray-300"
                    onClick={() => rate("dislike")}
                >
                    {film.dislikes}

                    <ThumbsDown strokeWidth={2} />
                </Button>
            </div>
        );
    };

    return (
        <>
            <Dialog>
                <div className="relative overflow-hidden transition-transform duration-300 ease-in-out rounded-lg shadow-lg group hover:shadow-xl hover:-translate-y-2">
                    <div className="relative">
                        <DialogTrigger asChild>
                            <img
                                src={film.posterPreUrl}
                                alt={film.title}
                                width={500}
                                height={750}
                                className="object-cover w-full h-[300px]"
                                style={{
                                    aspectRatio: "500/750",
                                    objectFit: "cover",
                                }}
                            />
                        </DialogTrigger>
                        <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2     opacity-0 group-hover:opacity-100 transition-opacity duration-300 pointer-events-none group-hover:pointer-events-auto ">
                            <Thumbs hoverCard={true} />
                        </div>
                    </div>
                    <DialogTrigger asChild>
                        <div className="flex flex-col p-4 bg-background w-full gap-1">
                            <h3 className="text-xl font-bold truncate">
                                {film.title}
                            </h3>
                            <div className="flex w-full flex-row gap-1 overflow-x-scroll scrollbar-hide">
                                {film.tags.map((tag, index) => (
                                    <Badge
                                        key={index}
                                        className="rounded-full text-nowrap font-medium"
                                    >
                                        {tag.name}
                                    </Badge>
                                ))}
                            </div>
                        </div>
                    </DialogTrigger>
                </div>
                <DialogContent
                    style={{ backgroundImage: `url(${film.backdropUrl})` }}
                    className="bg-cover bg-center  min-w-[95%] min-h-[70%]  border-none"
                >
                    <div className="absolute inset-0 bg-black/50 rounded-lg z-[-1000]" />

                    <div className="w-full h-full z-10 relative flex flex-col  justify-between">
                        <div className="flex flex-col gap-3">
                            <DialogTitle className="text-7xl text-white">
                                {film.title}
                            </DialogTitle>
                            <div className="flex w-full  flex-row gap-1 overflow-x-scroll scrollbar-hide gap-4  ">
                                {film.tags.map((tag) => (
                                    <a className=" text-primary-foreground text-xs font-normal rounded-full text-nowrap  bg-opacity-50">
                                        {tag.name}
                                    </a>
                                ))}
                            </div>
                        </div>
                        <DialogDescription className="flex text-white flex-col gap-2 items-end">
                            <div className="pr-4">
                                <Thumbs hoverCard={false} />
                            </div>
                            {film.description}
                        </DialogDescription>
                    </div>
                </DialogContent>
            </Dialog>
            <Dialog>
                <DialogContent className="">sdhhfhfh</DialogContent>
            </Dialog>
        </>
    );
};
