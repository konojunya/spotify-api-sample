<template>
  <section>
    <div class="wrapper">
      <figure class="image">
        <img :src="artwork" alt="track artwork">
      </figure>
      <div class="content">
        <p>{{item.track.name}}</p>
        <p>{{item.track.album.name}}</p>
        <p>by: {{artists}}</p>
        <button @click="select">{{item.isPlay ? "停止" : "再生"}}</button>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  props: ["item", "playOrStop"],
  computed: {
    artists() {
      return this.item.track.artists.map((artist) => artist.name).join(", ")
    },
    artwork() {
      return this.item.track.album.images[0].url
    }
  },
  methods: {
    select() {
      this.playOrStop(this.item.track.id)
    }
  }
}
</script>

<style lang="scss" scoped>
.wrapper {
  display: flex;
}
.image {
  width: 30%;
  img {
    display: block;
    width: 100%;
  }
}
</style>
