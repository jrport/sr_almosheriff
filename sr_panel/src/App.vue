<script setup>
	import SideBar from './components/SideBar.vue'
	import SearchBar from './components/SearchBar.vue'
	import { RouterView } from 'vue-router';
</script>

<template>
	<div class="h-screen w-screen flex flex-row">
		<SideBar/>
		<div class="flex flex-col w-full text-black">
			<SearchBar/>
			<RouterView />
		</div>
	</div>
</template>
