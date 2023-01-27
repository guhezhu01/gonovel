<template>
    <div class="vue-text" ref="text" :style="textStyle">
        {{value.slice(0,40) + '...'}}
    </div>
</template>
  
<script>
export default {
    // 显示文字组件，可以设置最多显示几行，超过后会隐藏并且鼠标hover显示全部信息（需要给组件设置宽度）
    name: "VueText",
    props: {
        value: {
            type: String,
            default: "",
        },
        row: {
            //最多显示几行，超过后会...隐藏 为0时不隐藏
            type: [Number, String],
            default: 0,
        },
    },
    computed: {
        text() {
            return this.$refs.text;
        },
    },
    data() {
        return {
            isShowHover: false,
            textStyle: {},
            div: null,
        };
    },
    watch: {
        row: function (val) {
            this.init();
        },
        value: function () {
            this.isShowHover = false;
            this.textStyle = {
                cursor: "text",
            };
            this.$nextTick(() => {
                this.getStyle(this.row - 0);
            });
        },
    },
    mounted() {
        this.init();
    },
    methods: {
        init() {
            this.div = document.querySelector(".hover-wrap");
            if (!this.div && this.row - 0) {
                this.div = document.createElement("div");
                this.div.className = "hover-wrap";
                this.div.style.top = 0;
                this.div.style.left = 0;
                document.body.append(this.div);
            }
            if (this.div && this.row - 0) {
                this.getStyle(this.row - 0);
            }
        },
        getStyle(val) {
            let lineHeight = getComputedStyle(this.text).lineHeight.replace("px", "") - 0 || 20;
            let height = getComputedStyle(this.text).height.replace("px", "") - 0;
            if (height > lineHeight * val) {
                this.isShowHover = true;
                this.textStyle = {
                    height: `${lineHeight * val}px`,
                    overflow: "hidden",
                    textOverflow: "ellipsis",
                    display: "-webkit-box",
                    webkitLineClamp: val,
                    webkitBoxOrient: "vertical",
                    cursor: "pointer",
                };
            } else {
                this.isShowHover = false;
                this.textStyle = {
                    cursor: "text",
                };
            }
        },
  
  
    },
};
</script>
  
<style lang="css">
.hover-wrap {
    position: absolute;
    background-color: #fff;
    color: #666;
    font-size: 14px;
    line-height: 1.4;
    padding: 6px 10px;
    border-radius: 5px;
    border: 1px solid #eee;
    max-width: 400px;
    transition: opacity 0.3s linear;
    z-index: -1;
    opacity: 0;
}

.hover-wrap.active {
    z-index: 1999;
    opacity: 1;
}
</style>