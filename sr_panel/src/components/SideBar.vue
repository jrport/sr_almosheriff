<script setup>
import { ref } from 'vue'
import { RouterLink } from 'vue-router'
</script>

<template>
	<nav class="side_bar w-2/12 flex text-white flex-col gap-5 items-center py-5">
		<img src="../assets/logo.png" alt="Logo do São Roque" class="w-10/12">
		<menu class="flex-col flex items-center  justify-center">
			<RouterLink to="/pedidos" class="hover:scale-110 hover:text-stone-300 duration-150">Pedidos</RouterLink>
			<RouterLink to="/pedidos" class="hover:scale-110 hover:text-stone-300 duration-150">Compras</RouterLink>
			<RouterLink to="/pedidos" class="hover:scale-110 hover:text-stone-300 duration-150">Inventários</RouterLink>
			<RouterLink to="/pedidos" class="hover:scale-110 hover:text-stone-300 duration-150">Performances</RouterLink>
		</menu>
		<small class="absolute bottom-2 text-white">V-1.0</small>
	</nav>
</template>

<style scoped>
.side_bar{
	background-color: #00063D;
}
</style>
