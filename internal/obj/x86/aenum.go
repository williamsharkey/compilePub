// Code generated by x86avxgen. DO NOT EDIT.

package x86

import "github.com/williamsharkey/compilePub/internal/obj"

//go:generate go run ../stringer.go -i $GOFILE -o anames.go -p x86

const (
	AAAA = obj.ABaseAMD64 + obj.A_ARCHSPECIFIC + iota
	AAAD
	AAAM
	AAAS
	AADCB
	AADCL
	AADCQ
	AADCW
	AADCXL
	AADCXQ
	AADDB
	AADDL
	AADDPD
	AADDPS
	AADDQ
	AADDSD
	AADDSS
	AADDSUBPD
	AADDSUBPS
	AADDW
	AADJSP
	AADOXL
	AADOXQ
	AAESDEC
	AAESDECLAST
	AAESENC
	AAESENCLAST
	AAESIMC
	AAESKEYGENASSIST
	AANDB
	AANDL
	AANDNL
	AANDNPD
	AANDNPS
	AANDNQ
	AANDPD
	AANDPS
	AANDQ
	AANDW
	AARPL
	ABEXTRL
	ABEXTRQ
	ABLENDPD
	ABLENDPS
	ABLSIL
	ABLSIQ
	ABLSMSKL
	ABLSMSKQ
	ABLSRL
	ABLSRQ
	ABOUNDL
	ABOUNDW
	ABSFL
	ABSFQ
	ABSFW
	ABSRL
	ABSRQ
	ABSRW
	ABSWAPL
	ABSWAPQ
	ABTCL
	ABTCQ
	ABTCW
	ABTL
	ABTQ
	ABTRL
	ABTRQ
	ABTRW
	ABTSL
	ABTSQ
	ABTSW
	ABTW
	ABYTE
	ABZHIL
	ABZHIQ
	ACDQ
	ACLC
	ACLD
	ACLFLUSH
	ACLI
	ACLTS
	ACMC
	ACMOVLCC
	ACMOVLCS
	ACMOVLEQ
	ACMOVLGE
	ACMOVLGT
	ACMOVLHI
	ACMOVLLE
	ACMOVLLS
	ACMOVLLT
	ACMOVLMI
	ACMOVLNE
	ACMOVLOC
	ACMOVLOS
	ACMOVLPC
	ACMOVLPL
	ACMOVLPS
	ACMOVQCC
	ACMOVQCS
	ACMOVQEQ
	ACMOVQGE
	ACMOVQGT
	ACMOVQHI
	ACMOVQLE
	ACMOVQLS
	ACMOVQLT
	ACMOVQMI
	ACMOVQNE
	ACMOVQOC
	ACMOVQOS
	ACMOVQPC
	ACMOVQPL
	ACMOVQPS
	ACMOVWCC
	ACMOVWCS
	ACMOVWEQ
	ACMOVWGE
	ACMOVWGT
	ACMOVWHI
	ACMOVWLE
	ACMOVWLS
	ACMOVWLT
	ACMOVWMI
	ACMOVWNE
	ACMOVWOC
	ACMOVWOS
	ACMOVWPC
	ACMOVWPL
	ACMOVWPS
	ACMPB
	ACMPL
	ACMPPD
	ACMPPS
	ACMPQ
	ACMPSB
	ACMPSD
	ACMPSL
	ACMPSQ
	ACMPSS
	ACMPSW
	ACMPW
	ACMPXCHG8B
	ACMPXCHGB
	ACMPXCHGL
	ACMPXCHGQ
	ACMPXCHGW
	ACOMISD
	ACOMISS
	ACPUID
	ACQO
	ACRC32B
	ACRC32Q
	ACVTPD2PL
	ACVTPD2PS
	ACVTPL2PD
	ACVTPL2PS
	ACVTPS2PD
	ACVTPS2PL
	ACVTSD2SL
	ACVTSD2SQ
	ACVTSD2SS
	ACVTSL2SD
	ACVTSL2SS
	ACVTSQ2SD
	ACVTSQ2SS
	ACVTSS2SD
	ACVTSS2SL
	ACVTSS2SQ
	ACVTTPD2PL
	ACVTTPS2PL
	ACVTTSD2SL
	ACVTTSD2SQ
	ACVTTSS2SL
	ACVTTSS2SQ
	ACWD
	ADAA
	ADAS
	ADECB
	ADECL
	ADECQ
	ADECW
	ADIVB
	ADIVL
	ADIVPD
	ADIVPS
	ADIVQ
	ADIVSD
	ADIVSS
	ADIVW
	ADPPD
	ADPPS
	AEMMS
	AENTER
	AEXTRACTPS
	AF2XM1
	AFABS
	AFADDD
	AFADDDP
	AFADDF
	AFADDL
	AFADDW
	AFCHS
	AFCLEX
	AFCMOVCC
	AFCMOVCS
	AFCMOVEQ
	AFCMOVHI
	AFCMOVLS
	AFCMOVNE
	AFCMOVNU
	AFCMOVUN
	AFCOMD
	AFCOMDP
	AFCOMDPP
	AFCOMF
	AFCOMFP
	AFCOMI
	AFCOMIP
	AFCOML
	AFCOMLP
	AFCOMW
	AFCOMWP
	AFCOS
	AFDECSTP
	AFDIVD
	AFDIVDP
	AFDIVF
	AFDIVL
	AFDIVRD
	AFDIVRDP
	AFDIVRF
	AFDIVRL
	AFDIVRW
	AFDIVW
	AFFREE
	AFINCSTP
	AFINIT
	AFLD1
	AFLDCW
	AFLDENV
	AFLDL2E
	AFLDL2T
	AFLDLG2
	AFLDLN2
	AFLDPI
	AFLDZ
	AFMOVB
	AFMOVBP
	AFMOVD
	AFMOVDP
	AFMOVF
	AFMOVFP
	AFMOVL
	AFMOVLP
	AFMOVV
	AFMOVVP
	AFMOVW
	AFMOVWP
	AFMOVX
	AFMOVXP
	AFMULD
	AFMULDP
	AFMULF
	AFMULL
	AFMULW
	AFNOP
	AFPATAN
	AFPREM
	AFPREM1
	AFPTAN
	AFRNDINT
	AFRSTOR
	AFSAVE
	AFSCALE
	AFSIN
	AFSINCOS
	AFSQRT
	AFSTCW
	AFSTENV
	AFSTSW
	AFSUBD
	AFSUBDP
	AFSUBF
	AFSUBL
	AFSUBRD
	AFSUBRDP
	AFSUBRF
	AFSUBRL
	AFSUBRW
	AFSUBW
	AFTST
	AFUCOM
	AFUCOMI
	AFUCOMIP
	AFUCOMP
	AFUCOMPP
	AFXAM
	AFXCHD
	AFXRSTOR
	AFXRSTOR64
	AFXSAVE
	AFXSAVE64
	AFXTRACT
	AFYL2X
	AFYL2XP1
	AHADDPD
	AHADDPS
	AHLT
	AHSUBPD
	AHSUBPS
	AIDIVB
	AIDIVL
	AIDIVQ
	AIDIVW
	AIMUL3Q
	AIMULB
	AIMULL
	AIMULQ
	AIMULW
	AINB
	AINCB
	AINCL
	AINCQ
	AINCW
	AINL
	AINSB
	AINSERTPS
	AINSL
	AINSW
	AINT
	AINTO
	AINVD
	AINVLPG
	AINW
	AIRETL
	AIRETQ
	AIRETW
	AJCC // >= unsigned
	AJCS // < unsigned
	AJCXZL
	AJCXZQ
	AJCXZW
	AJEQ // == (zero)
	AJGE // >= signed
	AJGT // > signed
	AJHI // > unsigned
	AJLE // <= signed
	AJLS // <= unsigned
	AJLT // < signed
	AJMI // sign bit set (negative)
	AJNE // != (nonzero)
	AJOC // overflow clear
	AJOS // overflow set
	AJPC // parity clear
	AJPL // sign bit clear (positive)
	AJPS // parity set
	ALAHF
	ALARL
	ALARW
	ALDDQU
	ALDMXCSR
	ALEAL
	ALEAQ
	ALEAVEL
	ALEAVEQ
	ALEAVEW
	ALEAW
	ALFENCE
	ALOCK
	ALODSB
	ALODSL
	ALODSQ
	ALODSW
	ALONG
	ALOOP
	ALOOPEQ
	ALOOPNE
	ALSLL
	ALSLW
	AMASKMOVOU
	AMASKMOVQ
	AMAXPD
	AMAXPS
	AMAXSD
	AMAXSS
	AMFENCE
	AMINPD
	AMINPS
	AMINSD
	AMINSS
	AMOVAPD
	AMOVAPS
	AMOVB
	AMOVBLSX
	AMOVBLZX
	AMOVBQSX
	AMOVBQZX
	AMOVBWSX
	AMOVBWZX
	AMOVDDUP
	AMOVHLPS
	AMOVHPD
	AMOVHPS
	AMOVL
	AMOVLHPS
	AMOVLPD
	AMOVLPS
	AMOVLQSX
	AMOVLQZX
	AMOVMSKPD
	AMOVMSKPS
	AMOVNTDQA
	AMOVNTIL
	AMOVNTIQ
	AMOVNTO
	AMOVNTPD
	AMOVNTPS
	AMOVNTQ
	AMOVO
	AMOVOU
	AMOVQ
	AMOVQL
	AMOVQOZX
	AMOVSB
	AMOVSD
	AMOVSHDUP
	AMOVSL
	AMOVSLDUP
	AMOVSQ
	AMOVSS
	AMOVSW
	AMOVUPD
	AMOVUPS
	AMOVW
	AMOVWLSX
	AMOVWLZX
	AMOVWQSX
	AMOVWQZX
	AMPSADBW
	AMULB
	AMULL
	AMULPD
	AMULPS
	AMULQ
	AMULSD
	AMULSS
	AMULW
	AMULXL
	AMULXQ
	ANEGB
	ANEGL
	ANEGQ
	ANEGW
	ANOTB
	ANOTL
	ANOTQ
	ANOTW
	AORB
	AORL
	AORPD
	AORPS
	AORQ
	AORW
	AOUTB
	AOUTL
	AOUTSB
	AOUTSL
	AOUTSW
	AOUTW
	APABSB
	APABSD
	APABSW
	APACKSSLW
	APACKSSWB
	APACKUSDW
	APACKUSWB
	APADDB
	APADDL
	APADDQ
	APADDSB
	APADDSW
	APADDUSB
	APADDUSW
	APADDW
	APALIGNR
	APAND
	APANDN
	APAUSE
	APAVGB
	APAVGW
	APBLENDW
	APCLMULQDQ
	APCMPEQB
	APCMPEQL
	APCMPEQQ
	APCMPEQW
	APCMPESTRI
	APCMPESTRM
	APCMPGTB
	APCMPGTL
	APCMPGTQ
	APCMPGTW
	APCMPISTRI
	APCMPISTRM
	APDEPL
	APDEPQ
	APEXTL
	APEXTQ
	APEXTRB
	APEXTRD
	APEXTRQ
	APEXTRW
	APHADDD
	APHADDSW
	APHADDW
	APHMINPOSUW
	APHSUBD
	APHSUBSW
	APHSUBW
	APINSRB
	APINSRD
	APINSRQ
	APINSRW
	APMADDUBSW
	APMADDWL
	APMAXSB
	APMAXSD
	APMAXSW
	APMAXUB
	APMAXUD
	APMAXUW
	APMINSB
	APMINSD
	APMINSW
	APMINUB
	APMINUD
	APMINUW
	APMOVMSKB
	APMOVSXBD
	APMOVSXBQ
	APMOVSXBW
	APMOVSXDQ
	APMOVSXWD
	APMOVSXWQ
	APMOVZXBD
	APMOVZXBQ
	APMOVZXBW
	APMOVZXDQ
	APMOVZXWD
	APMOVZXWQ
	APMULDQ
	APMULHRSW
	APMULHUW
	APMULHW
	APMULLD
	APMULLW
	APMULULQ
	APOPAL
	APOPAW
	APOPCNTL
	APOPCNTQ
	APOPCNTW
	APOPFL
	APOPFQ
	APOPFW
	APOPL
	APOPQ
	APOPW
	APOR
	APREFETCHNTA
	APREFETCHT0
	APREFETCHT1
	APREFETCHT2
	APSADBW
	APSHUFB
	APSHUFD
	APSHUFHW
	APSHUFL
	APSHUFLW
	APSHUFW
	APSIGNB
	APSIGND
	APSIGNW
	APSLLL
	APSLLO
	APSLLQ
	APSLLW
	APSRAL
	APSRAW
	APSRLL
	APSRLO
	APSRLQ
	APSRLW
	APSUBB
	APSUBL
	APSUBQ
	APSUBSB
	APSUBSW
	APSUBUSB
	APSUBUSW
	APSUBW
	APTEST
	APUNPCKHBW
	APUNPCKHLQ
	APUNPCKHQDQ
	APUNPCKHWL
	APUNPCKLBW
	APUNPCKLLQ
	APUNPCKLQDQ
	APUNPCKLWL
	APUSHAL
	APUSHAW
	APUSHFL
	APUSHFQ
	APUSHFW
	APUSHL
	APUSHQ
	APUSHW
	APXOR
	AQUAD
	ARCLB
	ARCLL
	ARCLQ
	ARCLW
	ARCPPS
	ARCPSS
	ARCRB
	ARCRL
	ARCRQ
	ARCRW
	ARDMSR
	ARDPMC
	ARDTSC
	AREP
	AREPN
	ARETFL
	ARETFQ
	ARETFW
	AROLB
	AROLL
	AROLQ
	AROLW
	ARORB
	ARORL
	ARORQ
	ARORW
	ARORXL
	ARORXQ
	AROUNDPD
	AROUNDPS
	AROUNDSD
	AROUNDSS
	ARSM
	ARSQRTPS
	ARSQRTSS
	ASAHF
	ASALB
	ASALL
	ASALQ
	ASALW
	ASARB
	ASARL
	ASARQ
	ASARW
	ASARXL
	ASARXQ
	ASBBB
	ASBBL
	ASBBQ
	ASBBW
	ASCASB
	ASCASL
	ASCASQ
	ASCASW
	ASETCC
	ASETCS
	ASETEQ
	ASETGE
	ASETGT
	ASETHI
	ASETLE
	ASETLS
	ASETLT
	ASETMI
	ASETNE
	ASETOC
	ASETOS
	ASETPC
	ASETPL
	ASETPS
	ASFENCE
	ASHLB
	ASHLL
	ASHLQ
	ASHLW
	ASHLXL
	ASHLXQ
	ASHRB
	ASHRL
	ASHRQ
	ASHRW
	ASHRXL
	ASHRXQ
	ASHUFPD
	ASHUFPS
	ASQRTPD
	ASQRTPS
	ASQRTSD
	ASQRTSS
	ASTC
	ASTD
	ASTI
	ASTMXCSR
	ASTOSB
	ASTOSL
	ASTOSQ
	ASTOSW
	ASUBB
	ASUBL
	ASUBPD
	ASUBPS
	ASUBQ
	ASUBSD
	ASUBSS
	ASUBW
	ASWAPGS
	ASYSCALL
	ASYSRET
	ATESTB
	ATESTL
	ATESTQ
	ATESTW
	AUCOMISD
	AUCOMISS
	AUNPCKHPD
	AUNPCKHPS
	AUNPCKLPD
	AUNPCKLPS
	AVADDPD
	AVADDPS
	AVADDSD
	AVADDSS
	AVADDSUBPD
	AVADDSUBPS
	AVAESDEC
	AVAESDECLAST
	AVAESENC
	AVAESENCLAST
	AVAESIMC
	AVAESKEYGENASSIST
	AVANDNPD
	AVANDNPS
	AVANDPD
	AVANDPS
	AVBLENDPD
	AVBLENDPS
	AVBLENDVPD
	AVBLENDVPS
	AVBROADCASTF128
	AVBROADCASTI128
	AVBROADCASTSD
	AVBROADCASTSS
	AVCMPPD
	AVCMPPS
	AVCMPSD
	AVCMPSS
	AVCOMISD
	AVCOMISS
	AVCVTDQ2PD
	AVCVTDQ2PS
	AVCVTPD2DQX
	AVCVTPD2DQY
	AVCVTPD2PSX
	AVCVTPD2PSY
	AVCVTPH2PS
	AVCVTPS2DQ
	AVCVTPS2PD
	AVCVTPS2PH
	AVCVTSD2SI
	AVCVTSD2SIQ
	AVCVTSD2SS
	AVCVTSI2SDL
	AVCVTSI2SDQ
	AVCVTSI2SSL
	AVCVTSI2SSQ
	AVCVTSS2SD
	AVCVTSS2SI
	AVCVTSS2SIQ
	AVCVTTPD2DQX
	AVCVTTPD2DQY
	AVCVTTPS2DQ
	AVCVTTSD2SI
	AVCVTTSD2SIQ
	AVCVTTSS2SI
	AVCVTTSS2SIQ
	AVDIVPD
	AVDIVPS
	AVDIVSD
	AVDIVSS
	AVDPPD
	AVDPPS
	AVERR
	AVERW
	AVEXTRACTF128
	AVEXTRACTI128
	AVEXTRACTPS
	AVFMADD132PD
	AVFMADD132PS
	AVFMADD132SD
	AVFMADD132SS
	AVFMADD213PD
	AVFMADD213PS
	AVFMADD213SD
	AVFMADD213SS
	AVFMADD231PD
	AVFMADD231PS
	AVFMADD231SD
	AVFMADD231SS
	AVFMADDSUB132PD
	AVFMADDSUB132PS
	AVFMADDSUB213PD
	AVFMADDSUB213PS
	AVFMADDSUB231PD
	AVFMADDSUB231PS
	AVFMSUB132PD
	AVFMSUB132PS
	AVFMSUB132SD
	AVFMSUB132SS
	AVFMSUB213PD
	AVFMSUB213PS
	AVFMSUB213SD
	AVFMSUB213SS
	AVFMSUB231PD
	AVFMSUB231PS
	AVFMSUB231SD
	AVFMSUB231SS
	AVFMSUBADD132PD
	AVFMSUBADD132PS
	AVFMSUBADD213PD
	AVFMSUBADD213PS
	AVFMSUBADD231PD
	AVFMSUBADD231PS
	AVFNMADD132PD
	AVFNMADD132PS
	AVFNMADD132SD
	AVFNMADD132SS
	AVFNMADD213PD
	AVFNMADD213PS
	AVFNMADD213SD
	AVFNMADD213SS
	AVFNMADD231PD
	AVFNMADD231PS
	AVFNMADD231SD
	AVFNMADD231SS
	AVFNMSUB132PD
	AVFNMSUB132PS
	AVFNMSUB132SD
	AVFNMSUB132SS
	AVFNMSUB213PD
	AVFNMSUB213PS
	AVFNMSUB213SD
	AVFNMSUB213SS
	AVFNMSUB231PD
	AVFNMSUB231PS
	AVFNMSUB231SD
	AVFNMSUB231SS
	AVGATHERDPD
	AVGATHERDPS
	AVGATHERQPD
	AVGATHERQPS
	AVHADDPD
	AVHADDPS
	AVHSUBPD
	AVHSUBPS
	AVINSERTF128
	AVINSERTI128
	AVINSERTPS
	AVLDDQU
	AVLDMXCSR
	AVMASKMOVDQU
	AVMASKMOVPD
	AVMASKMOVPS
	AVMAXPD
	AVMAXPS
	AVMAXSD
	AVMAXSS
	AVMINPD
	AVMINPS
	AVMINSD
	AVMINSS
	AVMOVAPD
	AVMOVAPS
	AVMOVD
	AVMOVDDUP
	AVMOVDQA
	AVMOVDQU
	AVMOVHLPS
	AVMOVHPD
	AVMOVHPS
	AVMOVLHPS
	AVMOVLPD
	AVMOVLPS
	AVMOVMSKPD
	AVMOVMSKPS
	AVMOVNTDQ
	AVMOVNTDQA
	AVMOVNTPD
	AVMOVNTPS
	AVMOVQ
	AVMOVSD
	AVMOVSHDUP
	AVMOVSLDUP
	AVMOVSS
	AVMOVUPD
	AVMOVUPS
	AVMPSADBW
	AVMULPD
	AVMULPS
	AVMULSD
	AVMULSS
	AVORPD
	AVORPS
	AVPABSB
	AVPABSD
	AVPABSW
	AVPACKSSDW
	AVPACKSSWB
	AVPACKUSDW
	AVPACKUSWB
	AVPADDB
	AVPADDD
	AVPADDQ
	AVPADDSB
	AVPADDSW
	AVPADDUSB
	AVPADDUSW
	AVPADDW
	AVPALIGNR
	AVPAND
	AVPANDN
	AVPAVGB
	AVPAVGW
	AVPBLENDD
	AVPBLENDVB
	AVPBLENDW
	AVPBROADCASTB
	AVPBROADCASTD
	AVPBROADCASTQ
	AVPBROADCASTW
	AVPCLMULQDQ
	AVPCMPEQB
	AVPCMPEQD
	AVPCMPEQQ
	AVPCMPEQW
	AVPCMPESTRI
	AVPCMPESTRM
	AVPCMPGTB
	AVPCMPGTD
	AVPCMPGTQ
	AVPCMPGTW
	AVPCMPISTRI
	AVPCMPISTRM
	AVPERM2F128
	AVPERM2I128
	AVPERMD
	AVPERMILPD
	AVPERMILPS
	AVPERMPD
	AVPERMPS
	AVPERMQ
	AVPEXTRB
	AVPEXTRD
	AVPEXTRQ
	AVPEXTRW
	AVPGATHERDD
	AVPGATHERDQ
	AVPGATHERQD
	AVPGATHERQQ
	AVPHADDD
	AVPHADDSW
	AVPHADDW
	AVPHMINPOSUW
	AVPHSUBD
	AVPHSUBSW
	AVPHSUBW
	AVPINSRB
	AVPINSRD
	AVPINSRQ
	AVPINSRW
	AVPMADDUBSW
	AVPMADDWD
	AVPMASKMOVD
	AVPMASKMOVQ
	AVPMAXSB
	AVPMAXSD
	AVPMAXSW
	AVPMAXUB
	AVPMAXUD
	AVPMAXUW
	AVPMINSB
	AVPMINSD
	AVPMINSW
	AVPMINUB
	AVPMINUD
	AVPMINUW
	AVPMOVMSKB
	AVPMOVSXBD
	AVPMOVSXBQ
	AVPMOVSXBW
	AVPMOVSXDQ
	AVPMOVSXWD
	AVPMOVSXWQ
	AVPMOVZXBD
	AVPMOVZXBQ
	AVPMOVZXBW
	AVPMOVZXDQ
	AVPMOVZXWD
	AVPMOVZXWQ
	AVPMULDQ
	AVPMULHRSW
	AVPMULHUW
	AVPMULHW
	AVPMULLD
	AVPMULLW
	AVPMULUDQ
	AVPOR
	AVPSADBW
	AVPSHUFB
	AVPSHUFD
	AVPSHUFHW
	AVPSHUFLW
	AVPSIGNB
	AVPSIGND
	AVPSIGNW
	AVPSLLD
	AVPSLLDQ
	AVPSLLQ
	AVPSLLVD
	AVPSLLVQ
	AVPSLLW
	AVPSRAD
	AVPSRAVD
	AVPSRAW
	AVPSRLD
	AVPSRLDQ
	AVPSRLQ
	AVPSRLVD
	AVPSRLVQ
	AVPSRLW
	AVPSUBB
	AVPSUBD
	AVPSUBQ
	AVPSUBSB
	AVPSUBSW
	AVPSUBUSB
	AVPSUBUSW
	AVPSUBW
	AVPTEST
	AVPUNPCKHBW
	AVPUNPCKHDQ
	AVPUNPCKHQDQ
	AVPUNPCKHWD
	AVPUNPCKLBW
	AVPUNPCKLDQ
	AVPUNPCKLQDQ
	AVPUNPCKLWD
	AVPXOR
	AVRCPPS
	AVRCPSS
	AVROUNDPD
	AVROUNDPS
	AVROUNDSD
	AVROUNDSS
	AVRSQRTPS
	AVRSQRTSS
	AVSHUFPD
	AVSHUFPS
	AVSQRTPD
	AVSQRTPS
	AVSQRTSD
	AVSQRTSS
	AVSTMXCSR
	AVSUBPD
	AVSUBPS
	AVSUBSD
	AVSUBSS
	AVTESTPD
	AVTESTPS
	AVUCOMISD
	AVUCOMISS
	AVUNPCKHPD
	AVUNPCKHPS
	AVUNPCKLPD
	AVUNPCKLPS
	AVXORPD
	AVXORPS
	AVZEROALL
	AVZEROUPPER
	AWAIT
	AWBINVD
	AWORD
	AWRMSR
	AXABORT
	AXACQUIRE
	AXADDB
	AXADDL
	AXADDQ
	AXADDW
	AXBEGIN
	AXCHGB
	AXCHGL
	AXCHGQ
	AXCHGW
	AXEND
	AXGETBV
	AXLAT
	AXORB
	AXORL
	AXORPD
	AXORPS
	AXORQ
	AXORW
	AXRELEASE
	AXTEST
	ALAST
)