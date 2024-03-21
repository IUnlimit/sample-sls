/// <reference types="vite/client" />

// 为所有的 Vue 单文件组件提供一个全局类型定义，避免 ts 类型报错
declare module '*.vue' {
    import { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
  }