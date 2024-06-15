<template>
  <el-table :data="grpServiceList" style="width: 100%">
    <el-table-column prop="grpcService" label="GrpcService" />
    <el-table-column prop="action" label="Action">
      <template #default="scope">
        <el-button type="primary" size="small" @click="showMethods(scope.row)" :disabled="disableButton(scope.row)">
          {{ scope.row.action }}
        </el-button>
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

const grpServiceList = ref([]);

onBeforeMount(async () => {
  try {
    const arr = await ListGrpcSymbol(endpoint, '');
    grpServiceList.value = arr.map((svc) => {
      return {
        grpcService: svc,
        action: 'showMethods',
      };
    });
    alert.setSuccessTitle(`connect with ${endpoint} successfully`);
  } catch (err) {
    console.error(err);
    alert.setErrorTitle(`fail to connect with ${endpoint}`);
  }
});

const showMethods = (row) => {
  const linkPath = `${path}/${row.grpcService}?endpoint=${endpoint}`;
  router.push(linkPath);
};

const disableButton = (row) => {
  const arr = [
    'grpc.reflection.v1.ServerReflection',
    'grpc.reflection.v1alpha.ServerReflection',
    'grpc.health.v1.Health',
  ];
  for (const svc of arr) {
    if (svc === row.grpcService) {
      return true;
    }
  }
  return false;
};
</script>

<style scoped></style>
