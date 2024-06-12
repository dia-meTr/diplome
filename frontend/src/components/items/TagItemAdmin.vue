<template>
    <div class="tag-info">
        <div class="tag-details">
            <h4>{{ tag.name }}</h4>
        </div>
        <div class="admin-buttons">
            <my-button @click="deleteTag">Видалити</my-button>
            
            <my-button @click="updateTag">Редагувати</my-button>
        </div>
    </div>

</template>


<script>
export default {
    name: 'TagItemAdmin',
    props: {
        tag: {
            type: Object,
            required: true
        }
    },
    methods: {
        deleteTag() {
            this.$emit('delete', this.tag.id)
        },
        updateTag() {
            this.$emit('update', this.tag)
        }
    }
}

</script>

<style scoped>

.tag-info {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
}

.tag-details {
    flex-grow: 1;
}

.admin-buttons {
    display: flex;
    flex-direction: column;
    gap: 5px;
}
</style>