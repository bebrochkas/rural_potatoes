import { Badge } from "@/components/ui/badge";
import ognivo from "@/assets/ognivo.png";

export const FilmCard = () => {
    return (
        <div className="relative overflow-hidden transition-transform duration-300 ease-in-out rounded-lg shadow-lg group hover:shadow-xl hover:-translate-y-2">
            <a href="#" className="absolute inset-0 z-10">
                <span className="sr-only">View</span>
            </a>
            <img
                src={ognivo}
                alt="Movie 8"
                width={500}
                height={750}
                className="object-cover w-full h-[300px]"
                style={{ aspectRatio: "500/750", objectFit: "cover" }}
            />
            <div className="flex flex-col p-4 bg-background w-full gap-1 ">
                <h3 className="text-xl font-bold">The Matrix</h3>
                <div className="flex w-full  flex-row gap-1 overflow-x-scroll scrollbar-hide ">
                    {[...Array(10)].map((x, i) => (
                        <Badge className="rounded-full text-nowrap">
                            Новый год
                        </Badge>
                    ))}
                </div>
            </div>
        </div>
    );
};
