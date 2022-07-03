<template>
    <v-app>
        <v-sheet>
            <v-container>
                <v-card v-if="sentences.length < LIMITS" class="pa-5">
                    <v-form @submit.prevent="handleSubmit">
                        <v-text-field v-model="inputSentence" variant="outlined" label="今日あったよかったことを3つ入力😄" />
                        <v-btn color="primary" type="submit">追加</v-btn>
                    </v-form>
                </v-card>
                <v-sheet class="text-center" v-else>
                    <h1>今日もいい1日でした！🎉</h1>
                </v-sheet>
            </v-container>
            <v-container>
                <v-row>
                    <v-col v-for="(sentence, index) in sentences" :key="sentence" cols="4">
                        <v-card height="150">
                            <v-card-title>
                                {{ index + 1 }}つ目
                                <v-spacer />
                                <v-icon @click="handleRemove(index)">mdi-delete</v-icon>
                            </v-card-title>
                            <v-card-text>{{ sentence }}</v-card-text>
                        </v-card>
                    </v-col>
                </v-row>
            </v-container>
        </v-sheet>
    </v-app>
</template>

<script setup lang="ts">
import '@mdi/font/css/materialdesignicons.css'
import { useSentences } from '~/composables/sentence'
const LIMITS = 3
const inputSentence = ref("")

const { sentences, createSentence, removeSentence } = useSentences()
const handleSubmit = () => {
    createSentence(inputSentence.value)
    inputSentence.value = ''
}

const handleRemove = (index: number) => {
    removeSentence(index)
}
</script>