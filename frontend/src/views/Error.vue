<template>
    <div class="container">
        <a-result status="404" title="404" sub-title="抱歉，您访问的页面不存在。">
            <template #extra>
                <a-button type="primary" @click="refresh">刷新试试</a-button>
            </template>
        </a-result>
    </div>
</template>

<script setup>
import router from '../router/index';

const refresh = () => {
    router.push('/')
}
</script>

<style scoped>
.container {
    width: 100%;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
}
</style>