<template>
  <el-table :data="grpMethodList" style="width: 100%">
    <el-table-column prop="grpcMethod" label="GrpcMethod" />
    <el-table-column prop="action" label="Action">
      <template #default="scope">
        <el-button type="primary" size="small" @click="invokeAPI(scope.row)">{{ scope.row.action }}</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup>
import { alertStore } from '@/stores';
import { ListGrpcSymbol } from '@wailsjs/go/backend/App';
import { onBeforeMount, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const router = useRouter();
const route = useRoute();
const alert = alertStore();
const { path } = route;
const { endpoint } = route.query;
const { svc } = route.params;

const grpMethodList = ref([]);

onBeforeMount(async () => {
  try {
    const arr = await ListGrpcSymbol(endpoint, svc);
    grpMethodList.value = arr.map((method) => {
      return {
        grpcMethod: method,
        action: 'invokeAPI',
      };
    });
    alert.setSuccessTitle(`succeed in listing grpc methods for ${svc}`);
  } catch (err) {
    console.error(err);
    alert.setErrorTitle(`fail to connect with ${endpoint}`);
  }
});

const invokeAPI = (row) => {
  const i = row.grpcMethod.lastIndexOf('.');
  const grpcFunc = row.grpcMethod.substring(i + 1);
  const linkPath = `${path}/${grpcFunc}?endpoint=${endpoint}`;
  router.push(linkPath);
};
</script>

<style scoped></style>
