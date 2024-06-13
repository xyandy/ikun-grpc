<template>
  <el-container>
    <el-header><MainHeader /></el-header>
    <el-main><RouterView /></el-main>
  </el-container>
</template>

<script setup>
import MainHeader from '@/components/MainHeader.vue';
import { RouterView } from 'vue-router';
</script>

<style scoped></style>
