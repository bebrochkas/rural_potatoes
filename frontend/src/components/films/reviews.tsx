import { axiosInst } from "@/api/axios";

import { ThumbsUp, ThumbsDown } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Review } from "@/interfaces/interfaces";
import {
    Drawer,
    DrawerContent,
    DrawerDescription,
    DrawerFooter,
    DrawerHeader,
    DrawerTitle,
} from "@/components/ui/drawer";
import { Textarea } from "@/components/ui/textarea";

import {
    Card,
    CardContent,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import React, { useEffect, useState } from "react";
import { createReview, howLongAgo } from "@/api/review";

interface reviewsDrawerProps {
    filmId: number;
    open: boolean;
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

export const ReviewsDrawer = ({
    open,
    setOpen,
    filmId,
}: reviewsDrawerProps) => {
    const [reviews, setReviews] = useState<Review[]>([]);
    const [positive, setPositive] = useState<boolean>(true);
    const [content, setContent] = useState<string>("");
    const fetchReviews = async () => {
        try {
            const response = await axiosInst.get<Review[]>(
                `reviews?id=${filmId}`,
            );
            setReviews(response.data);
        } catch (error) {
            console.log((error as Error).message);
        }
    };

    useEffect(() => {
        fetchReviews(); // Fetch reviews when the component mounts or filmId changes
    }, []);

    return (
        <Drawer open={open} onClose={() => setOpen(false)}>
            <DrawerContent>
                <DrawerHeader>
                    <DrawerTitle className="text-5xl">Комментарии</DrawerTitle>
                    <DrawerDescription>
                        Ознакомтесь с мнениями зрителей, оцените их или оставьте
                        свое собственное!
                    </DrawerDescription>
                </DrawerHeader>
                <div className="flex flex-row  justify-between grow-1 px-4 w-full h-[50vh] gap-4 mt-2">
                    {reviews.length > 0 && (
                        <div className={`flex flex-col w-full gap-2`}>
                            {reviews.map((review) => (
                                <Card
                                    className={`w-full h-fit gap-2   ${review.positive ? "bg-green-100" : "bg-red-100"}`}
                                >
                                    <CardHeader>
                                        <CardTitle className="font-normal">
                                            {review.user.username}
                                        </CardTitle>
                                    </CardHeader>
                                    <CardContent className="font-bold">
                                        <a>{review.content}</a>
                                    </CardContent>
                                    <CardFooter>
                                        <a className="text-sm text-gray-500">
                                            {howLongAgo(review.updated_at)}
                                        </a>
                                    </CardFooter>
                                </Card>
                            ))}
                        </div>
                    )}
                    <Card className="w-full h-fit">
                        <CardHeader>
                            <CardTitle>Ваш комментарий</CardTitle>
                        </CardHeader>
                        <CardContent>
                            <Textarea
                                value={content}
                                onChange={(value) => {
                                    setContent(value.target.value);
                                }}
                            />
                        </CardContent>
                        <CardFooter className={`flex flex-row gap-2 `}>
                            <div>
                                <Button
                                    variant={"ghost"}
                                    className={`rounded-full border-2 border-gray-300 ${positive && "bg-blue-500 text-white"} `}
                                    onClick={() => setPositive(true)}
                                >
                                    <ThumbsUp strokeWidth={2} />
                                </Button>
                                <Button
                                    variant="ghost"
                                    className={`rounded-full border-2 border-gray-300 ${!positive && "bg-blue-500 text-white"}`}
                                    onClick={() => setPositive(false)}
                                >
                                    <ThumbsDown strokeWidth={2} />
                                </Button>
                            </div>

                            <Button
                                onClick={async () => {
                                    await createReview(
                                        filmId,
                                        positive,
                                        content,
                                    );
                                    await fetchReviews();
                                }}
                            >
                                Отправить
                            </Button>
                        </CardFooter>
                    </Card>
                </div>
                <DrawerFooter></DrawerFooter>
            </DrawerContent>
        </Drawer>
    );
};
