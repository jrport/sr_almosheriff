<script setup>
	import SideBar from './components/SideBar.vue'
	import SearchBar from './components/SearchBar.vue'
</script>

<template>
	<div class="h-screen w-screen flex flex-row">
		<SideBar/>
		<div>
			<SearchBar/>
			<div>Inicial</div>
		</div>
	</div>
</template>
