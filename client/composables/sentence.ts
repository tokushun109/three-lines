export const useSentences = () => {
    const sentences = useState<string[]>('sentences', () => [])
    // よかったこと を追加するための関数
    const createSentence = (sentence: string) => {
        sentences.value.push(sentence)
    }
    // よかったこと を削除するための関数
    const removeSentence = (index: number) => {
        delete sentences.value[index]
    }

    return {
        sentences: readonly(sentences),
        createSentence,
        removeSentence
    }
}