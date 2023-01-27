<template>
  <HomeView></HomeView>
</template>

<script>
import HomeView from './views/HomeView.vue'
export default {
  // mounted() {
  //   window.addEventListener("unload", this.saveState)
  // },
  // methods: {
  //   saveState() {
  //     window.sessionStorage.setItem('state', JSON.stringify(this.$store.state))
  //   }
  // },
  components: {
    HomeView
  },
  name: "App"
}
</script>

