<template>
  <!-- <button @click="check">check</button> -->
  <el-row :gutter="24">
    <el-col :span="12">
      <!-- request -->
      <el-tabs v-model="requestArg.activeName">
        <el-tab-pane label="RequestPayload" name="Payload">
          <JsonEditorWrapper :options="requestOptions" v-model="requestArg.payloadJson" class="view" />
        </el-tab-pane>
        <el-tab-pane label="RequestMetadata" name="Metadata">
          <JsonEditorWrapper :options="requestOptions" v-model="requestArg.metadataJson" class="view" />
        </el-tab-pane>
      </el-tabs>

      <br />
      <el-button type="primary" @click="invokeAPI">InvokeAPI</el-button>
      <el-button :icon="Refresh" @click="reset">Reset</el-button>
      <el-button type="info" @click="showProto">ShowProto</el-button>

      <el-drawer v-model="drawer" title="Request/Response Proto Definition" size="50%" :with-header="true">
        <span class="name">{{ requestArg.name }}</span>
        <br />
        <br />
        <div v-html="requestArg.proto" />

        <span class="name">{{ responseArg.name }}</span>
        <br />
        <br />
        <div v-html="responseArg.proto" />
        Proto3 supports a canonical encoding in JSON, please check the below url:
        https://developers.google.com/protocol-buffers/docs/proto3#json
      </el-drawer>

      <!-- request -->
    </el-col>

    <el-col :span="12">
      <!-- response -->
      <el-tabs v-model="responseArg.activeName">
        <el-tab-pane label="ResponsePayload" name="Payload">
          <JsonEditorWrapper :options="responseOptions" v-model="responseArg.payloadJson" class="view" />
        </el-tab-pane>
        <el-tab-pane label="ResponseMetadata" name="Metadata">
          <JsonEditorWrapper :options="responseOptions" v-model="responseArg.metadataJson" class="view" />
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
import { onBeforeMount, reactive, ref } from 'vue';
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
const requestArg = reactive({
  name: '',
  proto: '',
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
const responseArg = reactive({
  name: '',
  proto: '',
  payloadJson: {},
  metadataJson: {},
  activeName: 'Payload',
});

function check() {
  console.log('check', requestArg.payloadJson, requestArg.metadataJson);
}

onBeforeMount(async () => {
  try {
    const methodMsg = await DescribeGrpcSymbol(endpoint, grpcMethod);
    await getRequestNameAndResponseName(methodMsg);

    const template = await GetGrpcMessageTemplate(endpoint, requestArg.name);
    requestArg.templateJson = JSON.parse(template);
    requestArg.payloadJson = requestArg.templateJson;
    alert.setSuccessTitle(`succeed in describing grpc symbol for ${grpcMethod}`);
  } catch (err) {
    console.error(err);
    alert.setErrorTitle(`fail to connect with ${endpoint}`);
  }
});

async function getRequestNameAndResponseName(str) {
  const reqBeg = str.indexOf('(');
  const reqEnd = str.indexOf(')');
  if (reqBeg < 0 || reqEnd < 0) {
    throw new Error(`fail to get request name, msg: ${str}`);
  }
  requestArg.name = str
    .slice(reqBeg + 1, reqEnd)
    .trim()
    .slice(1);

  const respBeg = str.lastIndexOf('(');
  const respEnd = str.lastIndexOf(')');
  if (respBeg < 0 || respEnd < 0) {
    throw new Error(`fail to get response name, msg: ${str}`);
  }
  responseArg.name = str
    .slice(respBeg + 1, respEnd)
    .trim()
    .slice(1);
  console.log(`str: ${str}, requestName: ${requestArg.name}, responseName: ${responseArg.name}`);

  requestArg.proto = await DescribeGrpcSymbol(endpoint, requestArg.name);
  requestArg.proto = requestArg.proto.replaceAll('{', '{<br/>').replaceAll(';', ';<br/>') + '<hr/>';
  responseArg.proto = await DescribeGrpcSymbol(endpoint, responseArg.name);
  responseArg.proto = responseArg.proto.replaceAll('{', '{<br/>').replaceAll(';', ';<br/>') + '<hr/>';
  console.log(requestArg.proto);
  console.log(responseArg.proto);
}

async function invokeAPI() {
  try {
    const result = await InvokeGrpc(endpoint, grpcMethod, requestArg.payloadJson, requestArg.metadataJson);
    console.log(result);
    responseArg.payloadJson = JSON.parse(result.RespBody);
    responseArg.metadataJson = result.RespHeader;
    alert.setSuccessTitle(`succeed in invoking API for ${grpcMethod}`);
  } catch (err) {
    console.error(err);
    alert.setErrorTitle(`fail to connect with ${endpoint}`);
  }
}

function reset() {
  requestArg.payloadJson = requestArg.templateJson;
  requestArg.metadataJson = { caller: 'ikun-grpc' };
}

const drawer = ref(false);
function showProto() {
  drawer.value = !drawer.value;
}
</script>

<style scoped>
.view {
  height: 70vh;
}

.name {
  color: blue;
}
</style>
