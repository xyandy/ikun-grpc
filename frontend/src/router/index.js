import InvokeGrpcMethod from '@/components/InvokeGrpcMethod.vue';
import ListGrpcFunction from '@/components/ListGrpcFunction.vue';
import ListGrpcService from '@/components/ListGrpcService.vue';
import MainContent from '@/components/MainContent.vue';
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: MainContent,
    },
    {
      path: '/grpc',
      component: ListGrpcService,
    },
    {
      path: '/grpc/:svc',
      component: ListGrpcFunction,
    },
    {
      path: '/grpc/:svc/:func',
      component: InvokeGrpcMethod,
    },
    {
      path: '/test',
      component: () => import('@/views/TestView.vue'),
    },
  ],
});

export default router;
