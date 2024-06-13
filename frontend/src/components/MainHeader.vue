<template>
  <div class="flex">
    <el-input v-model="endpoint" :placeholder="msg" clearable class="flex1" />
    <el-button type="primary" @click="handleConnect">Connect</el-button>
    <el-button type="success" @click="handleReset">Reset</el-button>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const endpoint = ref('');

const defaultEndpoint = '0.0.0.0:16565';
const msg = `Please enter {host}:{port}, otherwise use ${defaultEndpoint} by default`;

const handleConnect = () => {
  if (endpoint.value === '') {
    endpoint.value = defaultEndpoint;
  }
  router.push(`/grpc?endpoint=${endpoint.value}`);
};

const handleReset = () => {
  endpoint.value = '';
  router.push(`/`);
};
</script>

<style scoped>
.flex {
  display: flex;
  margin-top: 10px;
}

.flex1 {
  flex: 1;
  margin-right: 12px;
}
</style>
