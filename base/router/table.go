package router

type Table interface {
	Delete()
	Add()
	List()
	Insert()
}

/**

   input -p icmp -j reject  // input 应用层自己的数据
   output // 应用层自己的数据
   forward -p icmp -j accept   // 数据转发
   forward -s 192.168.. -j accpet
   forward  -i eth1 -j accept



 */
