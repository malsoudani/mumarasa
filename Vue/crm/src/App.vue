<template>
  <div id="app" class="container">
    <SideBar />
    <AppContent />
  </div>
</template>

<script>
  import AppContent from '@/components/AppContent'
  import SideBar from '@/components/SideBar'
  export default {
    name:"App",
    components: {
      AppContent,
      SideBar
    }
  }
</script>

<style lang="scss" scoped>
  @import "./assets/styles/page.scss";
</style>