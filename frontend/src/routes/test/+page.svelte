<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Button, Dropdown, DropdownItem, Checkbox, Search } from 'flowbite-svelte';
	import { ChevronDownSolid, UserRemoveSolid } from 'flowbite-svelte-icons';
  </script>
  
  <Button>Dropdown search<ChevronDownSolid class="w-3 h-3 ms-2 text-white dark:text-white" /></Button>
  <Dropdown class="overflow-y-auto px-3 pb-3 text-sm h-44">
	<div slot="header" class="p-3">
	  <Search size="md" />
	</div>
	<li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
	  <Checkbox>Robert Gouth</Checkbox>
	</li>
	<li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
	  <Checkbox>Jese Leos</Checkbox>
	</li>
	<li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
	  <Checkbox checked>Bonnie Green</Checkbox>
	</li>
	<li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
	  <Checkbox>Jese Leos</Checkbox>
	</li>
	<li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
	  <Checkbox>Robert Gouth</Checkbox>
	</li>
	<li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
	  <Checkbox>Bonnie Green</Checkbox>
	</li>
	<a slot="footer" href="/" class="flex items-center p-3 -mb-1 text-sm font-medium text-red-600 bg-gray-50 hover:bg-gray-100 dark:bg-gray-700 dark:hover:bg-gray-600 dark:text-red-500 hover:underline">
	  <UserRemoveSolid class="w-4 h-4 me-2 text-primary-700 dark:text-primary-700" />Delete user
	</a>
  </Dropdown>