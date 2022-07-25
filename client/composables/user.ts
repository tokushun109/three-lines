import { useFetch } from "@nuxt/bridge/dist/runtime"
import axios from 'axios'
import { ILoginForm } from "~/types/user"

export const useSentences = () => {
    // const sentences = useState<string[]>('sentences', () => [])
    // サインアップ
    const { fetch, fetchState } = useFetch(async (loginForm: ILoginForm) => {
        await axios.post('https://myapi.com/name', loginForm)
    })
    return {
        createSentence,
    }
}