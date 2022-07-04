<template>
    <v-layout>
        <v-app>
            <v-app-bar>
                <v-container>
                    <strong>ポジティブ3行日記</strong>
                </v-container>
            </v-app-bar>
            <v-main>
                <v-sheet height="93vh" color="grey-lighten-3">
                    <slot></slot>
                </v-sheet>
            </v-main>
        </v-app>
    </v-layout>
</template>