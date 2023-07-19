# Javascript

## 原型链
- 所有对象都继承空对象：null
- Object：原型对象
- 其他对象
js中所有的对象都有一个内置的属性，这个内置属性叫prototype（原型）。原型本身也是一个对象，原型对象也会有它自己的原型，逐渐构成了原型链，原型链终止与原型是null的对象。

## Promise 异步

### async 和 await 
在异步函数内部使用await，可以用同步的写法来写异步代码。async 标识的函数是异步方法，返回Promise对象。

## undefined 和  null 区别
undefined 是类型，null是值。一个函数如何没有使用return 返回值，那么它的返回值就是 undefined。可以使用三种方式判断是不是undefined：
- === 全等符判断是否等于值 undefined，示例 
    ```js
    var x 
    if (x === undefined) {
        // 执行这里
    }
    ```
- typeof 判断类型，示例：
    ```js
    var x 
    if (typeof x === 'undefined'){
        // 执行这里
    }
    ```
- 使用 void 操作符替代，示例：
    ```js
    var x 
    if (x === void 0){
        // 执行这里
    }
    if (y === void 0){
        // 不执行，因为 y 未定义
    }
    ```

## for ... in 和 for 的区别？
for in 会变量所有属性

# Vue 

