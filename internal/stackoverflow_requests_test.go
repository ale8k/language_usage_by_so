package internal

import (
	"reflect"
	"testing"
)

func TestGetCreatedQuestionsSync(t *testing.T) {
	type args struct {
		page              int
		pageSize          int
		fromDateInSeconds int
		tags              string
	}

	tests := []struct {
		name    string
		args    args
		want    *SOQuestionResp
		wantErr bool
	}{
		{
			// this is obviously a flakty test due to watching an id that can change
			// but wanted to use data table for learning purposes in Go
			name: "test it responds as expected with specific conditions",
			args: args{
				page:              1,
				pageSize:          1,
				fromDateInSeconds: 1637971200,
				tags:              "go",
			},
			want: &SOQuestionResp{
				Items: []SOQuestionItem{
					{
						QuestionID: 70137427,
					},
				},
			},
		},
		{
			name: "test it disables ridonkulous page sizes",
			args: args{
				page:              1,
				pageSize:          100000,
				fromDateInSeconds: 1637971200,
				tags:              "go",
			},
			want: &SOQuestionResp{
				Items: []SOQuestionItem{
					{
						QuestionID: 70137006,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCreatedQuestionsSync(tt.args.page, tt.args.pageSize, tt.args.fromDateInSeconds, tt.args.tags)

			if (err != nil) != tt.wantErr && got == nil {
				t.Errorf("GetCreatedQuestionsSync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && got.Items[0].QuestionID != tt.want.Items[0].QuestionID {
				t.Errorf("GetCreatedQuestionsSync() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_getUrlCopy(t *testing.T) {
	t.Run("test it returns a pointer deepcopy of URL",
		func(t *testing.T) {
			got := getUrlCopy(soUrl)

			if !reflect.DeepEqual(*got, soUrl) {
				t.Errorf("getUrlCopy() = %v, want %v", got, soUrl)
			}
		},
	)
}
