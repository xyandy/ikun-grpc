<template>
  <el-container>
    <el-header><MainHeader /></el-header>
    <el-main><RouterView :key="route.fullPath" class="route-view" /></el-main>
  </el-container>
</template>

<script setup>
import MainHeader from '@/components/MainHeader.vue';
import { RouterView, useRoute } from 'vue-router';

const route = useRoute();
</script>

<style scoped>
.route-view {
  margin-top: 25px;
}
</style>
