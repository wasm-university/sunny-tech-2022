
![bg](#3AF1F2)
![fg](#000000)

### Avant de faire du Go
# 👀 1er module Wasm en C

---
<!--
`main.c`
```c
#define WASM_EXPORT __attribute__((visibility("default")))

WASM_EXPORT
float power(float number, int pow) {
 float res = number;
   for (int i = 0;i < pow - 1; i++) {
     res = res * number;
   }
 return res;
}

WASM_EXPORT
char* greet()
{
    static char str[12] = "hello world!";
    return (char*)str;
}
```

---
#### Build

```bash
clang --target=wasm32 \
  --no-standard-libraries -Wl,--export-all -Wl, \
  --no-entry -o main.wasm main.c
```

---

`index.html`
```javascript
WebAssembly.instantiateStreaming(fetch("main.wasm"))
  .then(({ instance }) => {
    console.log("👋 main.wasm is loaded")

    const value = instance.exports.power(2, 2)

    console.log(`🤖 value: ${value}`)
    console.log(`👋 greet: ${instance.exports.greet()}`)

  })
  .catch(error => {
    console.log("😡 ouch", error)
  })
```

---
-->
![bg](#000000)
![fg](#FFFFFF)
# Démo 🚀


<a href="https://github.com/wasm-university/sunny-tech-2022/tree/main/00-c-web" target="_blank">00-c-web</a>

---
