<template>
  <div class="flex">
    <el-input v-model="endpoint" :placeholder="msg" clearable class="flex1" />
    <el-button type="primary" @click="handleConnect">Connect</el-button>
    <el-button type="info" @click="back">Back</el-button>
    <el-button type="success" @click="home">Home</el-button>
  </div>
  <el-alert v-if="alert.state.show" v-bind="alert.state" show-icon :closable="false" class="alert" />
</template>

<script setup>
import { alertStore } from '@/stores';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const endpoint = ref('');
const alert = alertStore();

const defaultEndpoint = '0.0.0.0:16565';
const msg = `Please enter {host}:{port}, default endpoint is ${defaultEndpoint}`;

const handleConnect = () => {
  if (endpoint.value === '') {
    endpoint.value = defaultEndpoint;
  }
  router.push(`/grpc?endpoint=${endpoint.value}`);
};

const home = () => {
  endpoint.value = '';
  alert.state.show = false;
  router.push(`/`);
};

const back = () => {
  router.back();
};
</script>

<style scoped>
.flex {
  display: flex;
  margin-top: 20px;
}

.flex1 {
  flex: 1;
  margin-right: 12px;
}

.alert {
  margin-top: 2px;
  width: 100%;
}
</style>
