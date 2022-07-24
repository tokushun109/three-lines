<template>
  <v-app>
    <v-app-bar fixed app>
      <v-container class="d-flex">
        <v-toolbar-title v-text="'ポジティブ3行日記'" />
        <v-spacer />
        <v-btn to="/">sign up</v-btn>
      </v-container>
    </v-app-bar>
    <v-main>
      <v-container color="secondary">
        <Nuxt />
      </v-container>
    </v-main>
    <v-footer app>
      <v-container>
        <span>&copy; {{ new Date().getFullYear() }}</span>
      </v-container>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: 'DefaultLayout',
}
</script>
