<template>
  <div id="app" class="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style lang="stylus">
.background-image
  background-image: url(./assets/city-cropped.jpeg);
  background-size: cover;
  -webkit-filter: blur(5px);
  filter: blur(5px);
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;

.app
  position: fixed;
  left: 0;
  top: 0;
  padding-left: 15%;
  padding-top: 10%;
  width: 75%;
</style>
