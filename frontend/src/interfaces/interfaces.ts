export interface Review {
    id: number;
    userId: number;
    filmId: number;
    positive: boolean;
    content: string;
    updated_at: string;
    user: User;
}

interface User {
    username: string;
}

export interface Tag {
    id: number;
    name: string;
    hex: string;
    type: string;
}

export interface Film {
    id: number;
    title: string;
    description: string;
    posterPreUrl: string;
    posterUrl: string;
    backdropUrl: string;
    rate: number;
    likes: number;
    dislikes: number;
    userPositive: boolean;

    tags: Tag[];
}
