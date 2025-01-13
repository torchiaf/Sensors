<script>
import { allHash } from '@shell/utils/promise';

const MODULE = 'sensors.io.module';
const DEVICE = 'sensors.io.device';
const JOB = 'sensors.io.job';

export default {
  components: {
  },

  async fetch() {
    const inStore = this.$store.getters['currentProduct'].inStore;

    const hash = {};

    if (this.$store.getters[`${ inStore }/schemaFor`](MODULE)) { 
      hash.modules = this.$store.dispatch(`${ inStore }/findAll`, { type: MODULE });
    }

    if (this.$store.getters[`${ inStore }/schemaFor`](DEVICE)) { 
      hash.devices = this.$store.dispatch(`${ inStore }/findAll`, { type: DEVICE });
    }

    if (this.$store.getters[`${ inStore }/schemaFor`](JOB)) { 
      hash.jobs = this.$store.dispatch(`${ inStore }/findAll`, { type: JOB });
    }

    await allHash(hash);
  },

  data() {
    return {};
  },

  computed: {
    modules() {
      return this.$store.getters['management/all'](MODULE);
    },
    devices() {
      return this.$store.getters['management/all'](DEVICE);
    },
    jobs() {
      return this.$store.getters['management/all'](JOB);
    }
  }
};
</script>

<template>

  <h1>Sensors</h1>

  <div
      v-for="(m, i) in modules"
  >
    modules: {{ m.name }}
  </div>
  <div
      v-for="(d, i) in devices"
  >
    devices: {{ d.name }}
  </div>
  <div
      v-for="(j, i) in jobs"
  >
    jobs: {{ j.name }}
  </div>
</template>

<style lang="scss" scoped>
</style>
