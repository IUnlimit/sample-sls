<template>
  <lay-layout class="sls-layout">
    <lay-header>Header</lay-header>
    <lay-body>
      <lay-layout>
        <lay-side><Side></Side></lay-side>
        <lay-body>Content</lay-body>
      </lay-layout>
    </lay-body>
    <lay-footer>Footer</lay-footer>
  </lay-layout>
 </template>

<script setup lang="ts">
    import Side from './components/Side.vue'
</script>

<style lang="less" scoped>
html,
body {
  height: 100%;
  margin: 0;
  padding: 0;
}

.layui-layout {
  height: 100%;
}

.sls-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh; /* 确保至少占据视口的高度 */

  .layui-footer,
  .layui-header {
    line-height: 60px;
    text-align: center;
    background: #87ca9a;
    color: white;
  }

  .layui-side {
    display: flex;
    background: #77c38c;
    align-items: center;
    justify-content: center;
    color: white;
  }

  .layui-body {
    display: flex;
    flex: 1;
    background: #5FB878;
    align-items: center;
    justify-content: center;
    color: white;
  }
}
</style>
