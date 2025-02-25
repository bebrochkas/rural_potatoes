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

    tags: Tag[];
}
