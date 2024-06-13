<script setup>
import MainView from '@/views/MainView.vue';
</script>

<template>
  <el-row>
    <el-col :span="1"></el-col>
    <el-col :span="22">
      <MainView />
    </el-col>
    <el-col :span="1"></el-col>
  </el-row>
</template>

<style scoped>
/* .debug {
  background-color: azure;
  border: 2px solid red;
  margin-top: 50px;
} */
</style>
