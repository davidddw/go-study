$(function () {
    /*
     luckyNum为每次抽几人
     luckyResult为抽奖结果的集合（数组）
     luckyNum为5那么luckyResult的length也为5
     */
    var Obj = {};
    Obj.luckyResult = [];
    Obj.luckyPrize = '';
    Obj.luckyNum = $(".select_lucky_number").val();
    /*
     一次抽几人改变事件
     */
    $(".select_lucky_number").bind('change', function () {
        Obj.luckyNum = $(this).val();
    })
    /*
     图片预加载
     */
    function loadImage(arr, callback) {
        var loadImageLen = 1;
        var arrLen = arr.length;
        $('.all_number').html("/" + arrLen);
        for (var i = 0; i < arrLen; i++) {
            var img = new Image(); //创建一个Image对象，实现图片的预下载
            img.onload = function () {
                img.onload = null;
                ++loadImageLen;
                $(".current_number").html(loadImageLen);
                if (loadImageLen == arrLen) {
                    callback(img); //所有图片加载成功回调；
                }
                ;
            }
            img.src = arr[i].image;
        }
    }

    /*
     把3D动画初始化，等待执行
     personArray为本地引入数据
     
    Obj.M = $('.container').lucky({
        row: 7, //每排显示个数  必须为奇数
        col: 5,//每列显示个数  必须为奇数
        depth: 5, //纵深度
        iconW: 30, //图片的宽
        iconH: 30, //图片的高
        iconRadius: 8, //图片的圆角
        data: personArray, //数据的地址数组
    });*/
    /*
    执行图片预加载并关闭加载试图
    
    loadImage(personArray, function (img) {
        $('.loader_file').hide();
    });*/
    /*
     若为ajax请求执行这段代码
     此为为ajax请求;
    */
    
     $.get('lucky/all',function(data){
         if(data){
             personArray = data; //此为数组

             //执行图片预加载并关闭加载试图
             loadImage(personArray, function (img) {
                $('.loader_file').hide();
             });
             Obj.M = $('.container').lucky({
	             row : 7, //每排显示个数  必须为奇数
	             col : 7, //每列显示个数  必须为奇数
	             depth : 6, //纵深度
	             iconW : 30, //图片的宽
	             iconH : 30, //图片的高
	             iconRadius : 8, //图片的圆角
	             data : personArray, //数据的地址数组
             });
         }
     });
     

    /*
     中奖人员展示效果
     传入当前中奖数组中单个的key
     */
    function showLuckyPeople(num) {
        setTimeout(function () {
            var $luckyEle = $('<img class="lucky_icon" />');
            var $userName = $('<p class="lucky_userName"></p>');
            var $fragEle = $('<div class="lucky_userInfo"></div>');
            $fragEle.append($luckyEle, $userName);
            $('.mask').append($fragEle);
            $(".mask").fadeIn(200);
            $luckyEle.attr('src', Obj.luckyResult[num].image);
            $userName.text(Obj.luckyResult[num].name)
            $fragEle.animate({
                'left': '50%',
                'top': '50%',
                'height': '200px',
                'width': '200px',
                'margin-left': '-100px',
                'margin-top': '-100px',
            }, 1000, function () {
                setTimeout(function () {
                    $fragEle.animate({
                        'height': '100px',
                        'width': '100px',
                        'margin-left': '100px',
                        'margin-top': '-50px',
                    }, 400, function () {
                        $(".mask").fadeOut(0);
                        $luckyEle.attr('class', 'lpl_userImage').attr('style', '');
                        $userName.attr('class', 'lpl_userName').attr('style', '');
                        $fragEle.attr('class', 'lpl_userInfo').attr('style', '');
                        $('.lpl_list.active').append($fragEle);
                    })
                }, 1000)
            })
        }, num * 2500)
        setTimeout(function () {
            $('.lucky_list').show();
        }, 2500)
    }

    /*
     停止按钮事件函数
     */
    $('#stop').click(function () {
        Obj.M.stop();
        $(".container").hide();
        $(this).hide();
        var i = 0;
        for (; i < Obj.luckyResult.length; i++) {
            showLuckyPeople(i);
        }
        $.get('lucky/totalNum', function(data){
        	if(data){
        		$(".lucky_number").html(data)
        	}
        })

    })
    
    /*
     开始按钮事件函数
     */
    $('#open').click(function () {
        $('.lucky_list').hide();
        $(".container").show();
        Obj.M.open();
        //此为ajax请求获奖结果
        $.get('lucky/luckyNum', {luckyNum: Obj.luckyNum,luckyPrize:Obj.luckyPrize},function(data){
        	if(data){
        		Obj.luckyResult = data;
                $("#stop").show(500);
        	}
        })
        //ajax获奖结果结束
        //此为人工写入获奖结果
        /*randomLuckyArr();
        setTimeout(function () {
            $("#stop").show(500);
        }, 1000)*/
        //人工获奖结果结束
    })

    /*
     前端写中奖随机数
     */
    function randomLuckyArr() {
        Obj.luckyResult = [];
        for (var i = 0; i < Obj.luckyNum; i++) {
            var random = Math.floor(Math.random() * personArray.length);
            if (Obj.luckyResult.indexOf(random) == -1) {
                Obj.luckyResult.push(random)
            } else {
                i--;
            }
        }
    }

    /*
     切换奖品代码块
     */
    function tabPrize() {
        var luckyDefalut = $(".lucky_prize_picture").attr('data-default');
        var index = luckyDefalut ? luckyDefalut : 1;
        tabSport(index);
        var lucky_prize_number = $('.lucky_prize_show').length;
        $('.lucky_prize_left').click(function () {
            $('.lucky_prize_right').addClass('active');
            index <= 1 ? 1 : --index;
            tabSport(index, lucky_prize_number);
        })
        $('.lucky_prize_right').click(function () {
            $('.lucky_prize_left').addClass('active');
            index >= lucky_prize_number ? lucky_prize_number : ++index;
            tabSport(index, lucky_prize_number);
        })

    }

    /*
     切换奖品左右按钮公共模块
     */
    function tabSport(i, lucky_prize_number) {
        if (i >= lucky_prize_number) {
            $('.lucky_prize_right').removeClass('active');
        }
        if (i <= 1) {
            $('.lucky_prize_left').removeClass('active');
        }
        Obj.luckyPrize = i;
        $('.lucky_prize_show').hide().eq(i - 1).show();
        $(".lucky_prize_title").html($('.lucky_prize_show').eq(i - 1).attr('alt'));
        $('.lpl_list').removeClass('active').hide().eq(i - 1).show().addClass('active');
    }
    tabPrize();
})