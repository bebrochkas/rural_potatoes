import { FilmCard } from "@/components/films/card";
export const HomePage = () => {
    return (
        <>
            <div className="container grid items-center gap-8  md:px-6  lg:gap-12 grid-cols-2 lg:grid-cols-4 md:grid-cols-3 ">
                {[...Array(10)].map((x, i) => (
                    <FilmCard />
                ))}
            </div>
        </>
    );
};
