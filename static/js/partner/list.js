define(function(require, exports, module) {


    exports.init = function(){
        init_style();
        select_state();
        add_partner();
    }

    function init_style(){
        $(".UL li[data-val='"+VALS.state+"']").addClass("select")
    }

    function select_state(){
        $(".UL li").click(function(){
            if($(this).hasClass("select")){
                return;
            }else{
                location.href = "/partner/list?state="+$(this).attr("data-val")
            }
        })
    }

    function add_partner(){
        $("#add_partner").click(function(){
            location.href = "/partner"
        })
    }

});
