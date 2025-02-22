import { Badge } from "@/components/ui/badge";
import { Film } from "@/interfaces/interfaces";

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";

export const FilmCard = ({ film }: { film: Film }) => {
    return (
        <Dialog>
            <DialogTrigger asChild>
                <div className="relative overflow-hidden transition-transform duration-300 ease-in-out rounded-lg shadow-lg group hover:shadow-xl hover:-translate-y-2">
                    <img
                        src={film.posterPreUrl}
                        alt={film.title}
                        width={500}
                        height={750}
                        className="object-cover w-full h-[300px]"
                        style={{ aspectRatio: "500/750", objectFit: "cover" }}
                    />
                    <div className="flex flex-col p-4 bg-background w-full gap-1 ">
                        <h3 className="text-xl font-bold truncate">
                            {film.title}
                        </h3>
                        <div className="flex w-full  flex-row gap-1 overflow-x-scroll scrollbar-hide ">
                            {film.tags.map((tag) => (
                                <Badge className="rounded-full text-nowrap  font-medium">
                                    {tag.name}
                                </Badge>
                            ))}
                        </div>
                    </div>
                </div>
            </DialogTrigger>
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
                        <div className="flex w-full  flex-row gap-1 overflow-x-scroll scrollbar-hide  ">
                            {film.tags.map((tag) => (
                                <Badge className="font-medium rounded-full text-nowrap  bg-opacity-50">
                                    {tag.name}
                                </Badge>
                            ))}
                        </div>
                    </div>
                    <DialogDescription className="text-white">
                        {film.description}
                    </DialogDescription>
                </div>
            </DialogContent>
        </Dialog>
    );
};
