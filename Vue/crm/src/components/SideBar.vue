<template>
    <div id="sidebar" class="flex-1">
        Reneder Contacts
        <List :items="contacts" itemKey="name" type="contact" />
    </div>
</template>

<script>
    import List from "@/components/general/List"
    export default {
        name: 'SideBar',
        components: {
            List
        },
        data() {
            return {
                contacts: [{name: "Hayder Alsoudani", id: 1}, {name: "Mohamed Alsoudani", id: 2}]
            };
        }
    }
</script>
<style scoped>

</style>