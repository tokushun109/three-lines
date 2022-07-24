import { readonly, useState } from "@nuxt/bridge/dist/runtime"

export const useSentences = () => {
    const sentences = useState<string[]>('sentences', () => [])
    // よかったこと を追加するための関数
    const createSentence = (sentence: string) => {
        sentences.value.push(sentence)
    }
    // よかったこと を削除するための関数
    const removeSentence = (index: number) => {
        sentences.value.splice(index, 1)
    }

    return {
        sentences: readonly(sentences),
        createSentence,
        removeSentence
    }
}