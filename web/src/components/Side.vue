<template>
  <!-- <lay-switch v-model="collapse"></lay-switch> -->
<lay-menu v-model:selected-key="selectedKey" tree="true" v-model:open-keys="openKeys6" :collapse="collapse">
    <lay-menu-item id="1">
        <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
    </lay-menu-item>
    <lay-menu-item id="2">
        <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
    </lay-menu-item>
    <lay-menu-item id="3">
        <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
    </lay-menu-item> 
    <lay-sub-menu title="目录" id="7">
        <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
        <lay-menu-item id="8">
                  <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
        </lay-menu-item> 
        <lay-menu-item id="9">
        <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
        </lay-menu-item>
        <lay-sub-menu id="10">
            <template #icon>
              <lay-icon type="layui-icon-home"></lay-icon> 
            </template>
            <template #title>
              首页
            </template>
            <lay-menu-item id="11">
                      <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
            </lay-menu-item> 
            <lay-menu-item id="12">
                      <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
            </lay-menu-item>
            <lay-sub-menu id="13">
              <template #icon>
                <lay-icon type="layui-icon-home"></lay-icon> 
              </template>
              <template #title>
              首页
              </template>
              <lay-menu-item id="14">
                        <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
              </lay-menu-item> 
              <lay-menu-item id="15">
                        <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
              </lay-menu-item>
              <lay-menu-item id="16">
                        <template #icon>
          <lay-icon type="layui-icon-home"></lay-icon> 
        </template>
        <template #title>
          首页
        </template>
              </lay-menu-item>
          </lay-sub-menu>
        </lay-sub-menu>
    </lay-sub-menu> 
  </lay-menu>
</template>

<script setup="ts">
import { ref } from 'vue'
    const selectedKey = ref("0")
    const openKeys6 = ref(["7"])
    const collapse = ref(false)   
</script>

<style scoped>
  .lay-menu {
    height: 100%;
  }
</style>