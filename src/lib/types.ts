export type Word = {
    strength_bars: number;
    infinitive: string;
    normalized_string: string;
    pos: string;
    last_practiced_ms: bigint;
    skill: string;
    related_lexemes: string[];
    last_practiced: string;
    strength: number;
    skill_url_title: string;
    gender: string;
    id: string;
    lexeme_id: string;
    word_string: string;
}

export type TenseMapping = {
    [key: string]: string;
}

export type TenseInfo = {
    tense_string: string;
    tense: string;
}

export type Tenses = {
   indicative: TenseInfo[];
   subjunctive: TenseInfo[];
   others: TenseInfo[];
}

export type PronounMapping = {
    tenses: TenseMapping;
    pronoun: string;
}

export type LanguageInfo = {
    pronoun_mapping: PronounMapping[];
    tenses: Tenses;
}

export interface User {
    id: string;
    name: string;
}