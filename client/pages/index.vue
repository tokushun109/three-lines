<template>
    <div>
        <v-container>
            <div class="py-2">
                <h2>{{ today }}</h2>
            </div>
            <v-card v-if="sentences.length < LIMITS" class="pa-5">
                <v-form @submit.prevent="handleSubmit">
                    <v-text-field v-model="inputSentence" variant="outlined" label="今日あったよかったことを3つ入力😄" />
                    <v-btn color="primary" type="submit">追加</v-btn>
                </v-form>
            </v-card>
            <v-card v-else height="150" class="text-center" color="green-lighten-4">
                <h1 class="fill-height" justify="center" style="line-height: 150px;">今日もいい1日でした！🎉</h1>
            </v-card>
        </v-container>
        <v-container v-if="sentences.length">
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
    </div>
</template>

<script setup lang="ts">
import { useSentences } from '~/composables/sentence'
import dayjs from 'dayjs';
import { ref } from 'vue';

const LIMITS = 3
const inputSentence = ref("")
const today = dayjs(new Date).format('YYYY年MM月DD日')

const { sentences, createSentence, removeSentence } = useSentences()
const handleSubmit = () => {
    createSentence(inputSentence.value)
    inputSentence.value = ''
}

const handleRemove = (index: number) => {
    removeSentence(index)
}
</script>
