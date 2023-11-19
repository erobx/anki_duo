import * as fs from 'fs';
import type { LanguageInfo, Word, PronounMapping, Tenses } from '../types';

interface JsonResponse {
    language_string: string;
    learning_language: string;
    from_language: string;
    language_information: LanguageInfo;
    vocab_overview: Word[];
}

interface Vocabulary {
    language_string: string;
    language_information: LanguageInfo;
    vocab_overview: Word[];
}

const jsonString = fs.readFileSync('vocabulary.json', 'utf-8');
const response: JsonResponse = JSON.parse(jsonString);

const language: string = response['language_string'];
const langInfo: LanguageInfo = response['language_information'];
const words: Word[] = response['vocab_overview'];

const vocab: Vocabulary = {
    language_string: language,
    language_information: langInfo,
    vocab_overview: words,
}

console.log(vocab)