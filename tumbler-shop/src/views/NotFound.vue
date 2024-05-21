<template>
  <div class="not-found">
    <div class="status">
      404
    </div>
    <div class="message">
      <h1>Page Not Found!</h1><br />
      <h2>Просимо вибачення за будь-які незручності, спричинені використанням нашого сервісу.</h2>
      <h3>Запитану веб-сторінку було перейменовано, наразі недоступна або видалено.</h3>
      <h3>Перевірте ще раз, чи правильна введена адреса.</h3>
    </div><br />
    <button @click="$router.push({name: 'mainpage'})">
      Перейти на головну сторінку
    </button>
  </div>
</template>


<style lang="scss" scoped>
.not-found {
  line-height: 1.2;
  text-align: center;
  padding: 8rem 2rem;
  .status {
    font-size: 16rem;
    color: $color-warning;
  }
  .message {
    font-size: 5rem;
  }
  button {
    height: 3rem;
    background-color: $color-warning;
    border: none;
    &:hover {
      cursor: pointer;
    }
  }
}
</style>
