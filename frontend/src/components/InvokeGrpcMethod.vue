<template>
  <!-- <button @click="check">check</button> -->
  <el-row :gutter="24">
    <el-col :span="12">
      <!-- request -->
      <el-tabs type="card" v-model="requestEditor.activeName">
        <el-tab-pane label="RequestPayload" name="Payload">
          <JsonEditorWrapper :options="requestOptions" v-model="requestEditor.payloadJson" class="view" />
        </el-tab-pane>
        <el-tab-pane label="RequestMetadata" name="Metadata">
          <JsonEditorWrapper :options="requestOptions" v-model="requestEditor.metadataJson" class="view" />
        </el-tab-pane>
      </el-tabs>

      <br />
      <el-button type="primary" @click="invokeAPI">InvokeAPI</el-button>
      <el-button :icon="Refresh" @click="reset">Reset</el-button>
      <!-- request -->
    </el-col>

    <el-col :span="12">
      <!-- response -->
      <el-tabs type="card" v-model="responseEditor.activeName">
        <el-tab-pane label="ResponsePayload" name="Payload">
          <JsonEditorWrapper :options="responseOptions" v-model="responseEditor.payloadJson" class="view" />
        </el-tab-pane>
        <el-tab-pane label="ResponseMetadata" name="Metadata">
          <JsonEditorWrapper :options="responseOptions" v-model="responseEditor.metadataJson" class="view" />
        </el-tab-pane>
      </el-tabs>
      <!-- response -->
    </el-col>
  </el-row>
</template>

<script setup>
import JsonEditorWrapper from '@/components/JsonEditorWrapper.vue';
import { alertStore } from '@/stores';
import { Refresh } from '@element-plus/icons-vue';
import { DescribeGrpcSymbol, GetGrpcMessageTemplate, InvokeGrpc } from '@wailsjs/go/backend/App';
import { onBeforeMount, reactive } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();
const alert = alertStore();

const { endpoint } = route.query;
const { svc, func } = route.params;
const grpcMethod = `${svc}.${func}`;

const requestOptions = {
  mode: 'code',
  modes: ['code', 'preview', 'view'],
};
const requestEditor = reactive({
  templateJson: {},
  payloadJson: {},
  metadataJson: { caller: 'ikun-grpc' },
  activeName: 'Payload',
});
const responseOptions = {
  mode: 'code',
  modes: [],
  onEditable: () => {
    return false;
  },
};
const responseEditor = reactive({
  payloadJson: {},
  metadataJson: {},
  activeName: 'Payload',
});

function check() {
  console.log('check', requestEditor.payloadJson, requestEditor.metadataJson);
}

onBeforeMount(async () => {
  try {
    const methodMsg = await DescribeGrpcSymbol(endpoint, grpcMethod);
    const beg = methodMsg.indexOf('(');
    const end = methodMsg.indexOf(')');
    if (beg < 0 || end < 0) {
      throw new Error(`fail to DescribeGrpcSymbol, endpoint: ${endpoint}, method: ${grpcMethod}, msg: ${methodMsg}`);
    }

    const requestName = methodMsg.slice(beg + 1, end).trim();
    const template = await GetGrpcMessageTemplate(endpoint, requestName);
    requestEditor.templateJson = JSON.parse(template);
    requestEditor.payloadJson = requestEditor.templateJson;
    alert.setSuccessTitle(`connect with ${endpoint} successfully`);
  } catch (err) {
    console.error(err);
    alert.setErrorTitle(`fail to connect with ${endpoint}`);
  }
});

async function invokeAPI() {
  try {
    const result = await InvokeGrpc(endpoint, grpcMethod, requestEditor.payloadJson, requestEditor.metadataJson);
    console.log(result);
    responseEditor.payloadJson = JSON.parse(result.RespBody);
    responseEditor.metadataJson = result.RespHeader;
    alert.setSuccessTitle(`connect with ${endpoint} successfully`);
  } catch (err) {
    console.error(err);
    alert.setErrorTitle(`fail to connect with ${endpoint}`);
  }
}

function reset() {
  requestEditor.payloadJson = requestEditor.templateJson;
  requestEditor.metadataJson = { caller: 'ikun-grpc' };
}
</script>

<style scoped>
.view {
  height: 70vh;
}
</style>
