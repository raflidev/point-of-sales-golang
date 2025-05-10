package service

import (
	"context"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/members/domain/entity"
	"golang-point-of-sales-system/modules/members/domain/repository"
	"golang-point-of-sales-system/modules/members/dto/request"
	"golang-point-of-sales-system/modules/members/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type MemberServiceImpl struct {
	MemberRepository repository.MemberRepository
	Validate         *validator.Validate
}

func NewMemberService(memberRepository repository.MemberRepository, validate *validator.Validate) MemberService {
	return &MemberServiceImpl{
		MemberRepository: memberRepository,
		Validate:         validate,
	}
}
func (service *MemberServiceImpl) Create(ctx context.Context, request request.MemberCreateRequest) response.MemberResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	member := entity.Member{
		Kode_member: request.Kode_member,
		Nama:        request.Nama,
		Telepon:     request.Telepon,
		Alamat:      request.Alamat,
		Keterangan:  request.Keterangan,
	}

	member = service.MemberRepository.Save(ctx, member)

	return helper.ToMemberResponse(member)
}

func (service *MemberServiceImpl) Update(ctx context.Context, request request.MemberUpdateRequest) response.MemberResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	member, err := service.MemberRepository.FindById(ctx, request.Id)
	if err != nil {
		panic(err)
	}

	member.Kode_member = request.Kode_member
	member.Nama = request.Nama
	member.Telepon = request.Telepon
	member.Alamat = request.Alamat
	member.Keterangan = request.Keterangan

	member, err = service.MemberRepository.Update(ctx, member)
	if err != nil {
		panic(err)
	}

	return helper.ToMemberResponse(member)
}

func (service *MemberServiceImpl) Delete(ctx context.Context, memberId uuid.UUID) {
	member, err := service.MemberRepository.FindById(ctx, memberId)
	if err != nil {
		panic(err)
	}

	service.MemberRepository.Delete(ctx, member)
}

func (service *MemberServiceImpl) FindById(ctx context.Context, memberId uuid.UUID) response.MemberResponse {
	member, err := service.MemberRepository.FindById(ctx, memberId)
	if err != nil {
		panic(err)
	}
	return helper.ToMemberResponse(member)
}

func (service *MemberServiceImpl) FindAll(ctx context.Context) []response.MemberResponse {
	members := service.MemberRepository.FindAll(ctx)
	var memberResponses []response.MemberResponse
	for _, member := range members {
		memberResponses = append(memberResponses, helper.ToMemberResponse(member))
	}
	return memberResponses
}
