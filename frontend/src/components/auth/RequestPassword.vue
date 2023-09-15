<template>
    <section>
        <h2>Forgot Password </h2>
    <form>
    <div class="user-box">
      <input type="text">
      <label>Email</label>
    </div>
    <a href="#" class="submit">Request Reset</a>
  </form>
  </section>
</template>

<script setup lang="ts">

</script>

<style scoped>
section {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 400px;
  padding: 40px;
  transform: translate(-50%, -50%);
  background: #101010;
  box-sizing: border-box;
  box-shadow: 0 15px 25px rgba(0,0,0,.6);
  border-radius: 10px;
}

section h2 {
  margin: 0 0 30px;
  padding: 0;
  color: #fff;
  text-align: center;
}



section .user-box {
  position: relative;
}

section .user-box input {
  width: 100%;
  padding: 10px 0;
  font-size: 16px;
  color: #fff;
  margin-bottom: 30px;
  border: none;
  border-bottom: 1px solid #fff;
  outline: none;
  background: transparent;
}
section .user-box label {
  position: absolute;
  top:0;
  left: 0;
  padding: 10px 0;
  font-size: 18px;
  color: #c972f6;
  pointer-events: none;
  transition: .5s;
}

section .user-box input:focus ~ label,
section .user-box input:valid ~ label {
  top: -20px;
  left: 0;
  color: #03e9f4;
  font-size: 14px;
}

section form .submit, .link{
  position: relative;
  display: inline-block;
  padding: 10px 20px;
  color: #03e9f4;
  font-size: 16px;
  text-decoration: none;
  text-transform: uppercase;
  overflow: hidden;
  transition: .5s;
  margin-top: 20px;
  border-radius: 5px;
  letter-spacing: 4px
}

section .submit:hover {
  background: #c972f6;
  color: #fff;
  border-radius: 5px;
  box-shadow: 0 0 5px #c972f6,
              0 0 25px #c972f6,
              0 0 25px #c972f6,
              0 0 25px #c972f6;
}
section a span {
  position: absolute;
  display: block;
}

.link {
    font-size: 12px;
}

</style>
