<template>
  <div id="app-container">
    <router-view v-slot="{ Component }">
      <Transition name="fade-slide" mode="out-in">
        <component :is="Component" />
      </Transition>
    </router-view>
  </div>
</template>