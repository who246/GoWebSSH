 var map = new Map();
   function close(id){ 
       var ws =  map.get(id);
	    if(ws != null)
	    ws.close();
		map.remove(id);
		return true;
	}
	 function openDilog(id,title,cmdId){
		var opt ={mask:false,height:350, close:close,param:id} 
		$.pdialog.open("/admin/service/trim?id="+id+(cmdId==null||cmdId==""?"":("&cmdId="+cmdId)),id, title,opt);　 
	}
 