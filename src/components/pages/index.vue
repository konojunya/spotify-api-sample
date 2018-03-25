<template>
  <section>
    <track-module :playOrStop="playOrStop" :item="item" v-for="item in items" :key="item.track.id"></track-module>
  </section>
</template>

<script>
import axios from 'axios'
import TrackModule from '../modules/track-module.vue'

export default {
  data() {
    return {
      items: [],
      someMusicPlaying: false,
      context: null,
      source: null
    }
  },
  created() {
    this.context = new AudioContext()
    axios.get("/api/tracks").then((res) => {
      console.log("Received")
      for(let item of res.data.items) {
        this.items.push({
          isPlay: false,
          track: item.track
        })
      }
    })
  },
  components: {
    TrackModule
  },
  methods: {
    playOrStop(id) {
      let stopDuration = 0
      if(this.someMusicPlaying) {
        this.stop()
      }
      for(let item of this.items) {
        if(id == item.track.id) {
          if(item.isPlay) {
            this.stop()
            item.isPlay = false
          } else {
            this.play(item.track.preview_url)
            stopDuration = item.track.duration_ms
            item.isPlay = true
          }
        } else {
          item.isPlay = false
        }
      }
      setTimeout(() => {
        console.log(stopDuration)
        for(let item of this.items) {
          if(id == item.id) {
            item.isPlay = false
          }
        }
        this.stop()
      }, stopDuration)
    },
    play(url) {
      this.someMusicPlaying = true
      this.source = this.context.createBufferSource()
      axios({ url, method: "GET", responseType: "arraybuffer" }).then((res) => {
        this.context.decodeAudioData(res.data, (buffer) => {
          this.source.buffer = buffer
          this.source.connect(this.context.destination)
          this.source.start(0)
        })
      })
    },
    stop() {
      this.someMusicPlaying = false
      this.source.stop(0)
    }
  }
}
</script>
