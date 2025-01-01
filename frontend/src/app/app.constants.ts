export const URL_CONSTANTS: { [key: string]: string } = {
    SEASONS: '/competitions/seasons',
    COMPETITIONS: '/competitions/seasons/:season',
    COMPETITION: '/competitions/:competitionId',
    RESULT_UPLOAD: '/result',
    COMPETITIONS_ALL: '/competitions',
    CATEGORIES: '/competitions/:competitionId/categories',
    CATEGORIES_ALL: '/competitions/categories',
    ROPE_CLIMBERS: '/rope-climbers',
    RESULT_LIST: '/result/:competitionId/:categoryId',
    TOP_RESULT_LIST: '/result/top/:categoryId',
    DELETE_COMPETITION: '/result/:competitionId',
    NOMINATION_CRITERIAS: '/nomination-criterias/seasons',
};

export const MEN_CATEGORY_ID = 3;