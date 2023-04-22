<template>
  <ul class="technologies">
    <li v-for="technology in technologies" v-bind:key="technology.name">
      <b>{{technology.name}}</b>: {{technology.details}}
    </li>
  </ul>
</template>

<script lang="ts">
import axios from 'axios'

let publicApiBaseUrl;

export default {
  name: 'TechItems',
  data() {
    return {
      technologies: []
    };
  },
  setup() {
    let runtimeConfig = useRuntimeConfig();
    publicApiBaseUrl = runtimeConfig.public.apiBaseUrl;
  },
  mounted() {
    console.log(`publicApiBaseUrl: ${publicApiBaseUrl}`)
    axios
      .get(`${publicApiBaseUrl}/api/technologies`)
      .then(response => (this.technologies = response.data));
  }
};
</script>

<style scoped>
.technologies {
  margin-top: 5px;
}
</style>
